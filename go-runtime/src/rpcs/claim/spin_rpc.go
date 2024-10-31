package rpcs_claim

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_player "cifarm-server/src/collections/player"
	collections_spin "cifarm-server/src/collections/spin"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type SpinRpcResponse struct {
	InventoryKey string `json:"inventoryKey"`
	SpinKey      string `json:"spinKey"`
}

func SpinRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//get reward tracker
	object, err := collections_player.ReadRewardTracker(ctx, logger, db, nk, collections_player.ReadRewardTrackerParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "reward tracker not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//check reward tracker
	rewardTracker, err := collections_common.ToValue[collections_player.RewardTracker](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	// get spin info
	object, err = collections_system.ReadSpinConfigure(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "spin configure not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	spinConfigure, err := collections_common.ToValue[collections_system.SpinConfigure](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	lastSpinDate := time.Unix(rewardTracker.SpinInfo.LastSpinTime, 0).UTC()
	nextSpinDate := lastSpinDate.Add(time.Duration(spinConfigure.FreeSpinTime) * time.Second)

	now := time.Now().UTC().Unix()

	result := now >= nextSpinDate.Unix()
	if !result {
		// you have use free spin today, now pay to spin
		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     userId,
			GoldAmount: -spinConfigure.SpinPrice,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
	}

	rewardTracker.SpinInfo.LastSpinTime = now
	rewardTracker.SpinInfo.SpinCount++

	//update the last claimed time
	err = collections_player.WriteRewardTracker(ctx, logger, db, nk, collections_player.WriteRewardTrackerParams{
		RewardTracker: *rewardTracker,
		UserId:        userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//spin time!
	//read spin
	objects, err := collections_spin.ReadMany(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	spins, err := collections_common.ToValues2[collections_spin.Spin](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var inventoryResult collections_inventories.WriteResult
	var spinResult collections_spin.Spin

	randomValue := rand.Float64()
	for _, spin := range spins {
		if randomValue < spin.ThresholdMax && randomValue >= spin.ThresholdMin {
			spinResult = *spin
			//take the spin and process the reward
			switch spin.Type {
			case collections_spin.TYPE_GOLD:
				{
					//add golds
					err := wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
						UserId:     userId,
						GoldAmount: -spin.GoldAmount,
						Metadata: map[string]interface{}{
							"name": "Spin reward",
							"time": time.Now().Format(time.RFC850),
						},
					})
					if err != nil {
						logger.Error(err.Error())
						return "", err
					}
					break
				}
			case collections_spin.TYPE_TOKEN:
				{
					err := wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
						UserId:      userId,
						TokenAmount: spin.TokenAmount,
						Metadata: map[string]interface{}{
							"name": "Spin reward",
							"time": time.Now().Format(time.RFC850),
						},
					})
					if err != nil {
						logger.Error(err.Error())
						return "", err
					}
					break
				}
			case collections_spin.TYPE_SUPPLY:
				{
					//add supplies
					_inventoryResult, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
						Inventory: collections_inventories.Inventory{
							ReferenceKey: spin.Key,
							Quantity:     spin.Quantity,
							Type:         collections_inventories.TYPE_SUPPLY,
							AsTool:       true,
						},
						UserId: userId,
					})
					if err != nil {
						logger.Error(err.Error())
						return "", err
					}
					inventoryResult = *_inventoryResult
					break
				}
			case collections_spin.TYPE_SEED:
				{
					//add supplies
					_inventoryResult, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
						Inventory: collections_inventories.Inventory{
							ReferenceKey: spin.Key,
							Quantity:     spin.Quantity,
							Type:         collections_inventories.TYPE_SEED,
						},
						UserId: userId,
					})
					if err != nil {
						logger.Error(err.Error())
						return "", err
					}
					inventoryResult = *_inventoryResult
					break
				}
			}
		}
	}

	_value, err := json.Marshal(SpinRpcResponse{
		InventoryKey: inventoryResult.Key,
		SpinKey:      spinResult.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
