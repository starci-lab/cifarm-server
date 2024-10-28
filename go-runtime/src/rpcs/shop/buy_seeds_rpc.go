package rpcs_shop

import (
	collections_common "cifarm-server/src/collections/common"
	collections_crops "cifarm-server/src/collections/crops"
	collections_inventories "cifarm-server/src/collections/inventories"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuySeedsRpcParams struct {
	Key      string `json:"key,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type BuySeedsRpcResponse struct {
	InventorySeedKey string `json:"inventorySeedKey,omitempty"`
}

func BuySeedsRpc(ctx context.Context,
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

	var params *BuySeedsRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_crops.ReadByKey(ctx, logger, db, nk, collections_crops.ReadByKeyParams{
		Key: params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	crop, err := collections_common.ToValue[collections_crops.Crop](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if crop == nil {
		errMsg := "crop not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !crop.AvailableInShop {
		errMsg := "crop not available in shop"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	totalCost := int64(crop.Price) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		GoldAmount: -totalCost,
		UserId:     userId,
		Metadata: map[string]interface{}{
			"name": "Buy seeds",
			"key":  params.Key,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	result, err := collections_inventories.Write(ctx,
		logger, db, nk,
		collections_inventories.WriteParams{
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

	value, err := json.Marshal(BuySeedsRpcResponse{
		InventorySeedKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
