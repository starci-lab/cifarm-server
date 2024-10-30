package rpcs_upgrades

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpgradeBuildingRpcParams struct {
	PlacedItemBuildingKey string `json:"placedItemBuildingKey"`
}

func UpgradeBuildingRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *UpgradeBuildingRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemBuildingKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "building not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	building, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if building.BuildingInfo.CurrentUpgrade == building.BuildingInfo.Building.MaxUpgrade {
		errMsg := "building already at max upgrade"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	//pay
	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		UserId:     userId,
		GoldAmount: -building.BuildingInfo.Building.Upgrades[building.BuildingInfo.CurrentUpgrade+1].UpgradePrice,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	//upgrade
	building.BuildingInfo.CurrentUpgrade++
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *building,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
