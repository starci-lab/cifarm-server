package rpcs_shop

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_seeds "cifarm-server/src/collections/seeds"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuySeedRpcParams struct {
	Key      string `json:"key"`
	Quantity int    `json:"quantity"`
}

type BuySeedRpcResponse struct {
	InventorySeedKey string `json:"inventorySeedKey"`
}

func BuySeedRpc(ctx context.Context,
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

	var params *BuySeedRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_seeds.ReadByKey(ctx, logger, db, nk, collections_seeds.ReadByKeyParams{
		Key: params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	seed, err := collections_common.ToValue[collections_seeds.Seed](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if seed == nil {
		errMsg := "seed not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	totalCost := int64(seed.Price) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		Amount: -totalCost,
		UserId: userId,
		Metadata: map[string]interface{}{
			"name": "Buy seeds",
			"key":  params.Key,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	result, err := collections_inventories.WriteSeed(ctx,
		logger, db, nk,
		collections_inventories.WriteSeedParams{
			Inventory: collections_inventories.Inventory{
				ReferenceKey: params.Key,
				Quantity:     params.Quantity,
				Type:         collections_inventories.TYPE_SEED,
			},
			UserId: userId,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuySeedRpcResponse{
		InventorySeedKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
