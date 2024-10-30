package rpcs_shop

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_supplies "cifarm-server/src/collections/supplies"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuySuppliesRpcParams struct {
	Key      string `json:"key"`
	Quantity int    `json:"quantity"`
}

type BuySuppliesRpcResponse struct {
	InventorySupplyKey string `json:"inventorySupplyKey"`
}

func BuySuppliesRpc(ctx context.Context,
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

	var params *BuySuppliesRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_supplies.ReadByKey(ctx, logger, db, nk, collections_supplies.ReadByKeyParams{
		Key: params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "supply not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	supply, err := collections_common.ToValue[collections_supplies.Supply](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if !supply.AvailableInShop {
		errMsg := "supply not available in shop"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	totalCost := int64(supply.Price) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		GoldAmount: -totalCost,
		UserId:     userId,
		Metadata: map[string]interface{}{
			"name": "Buy supplies",
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
				Type:         collections_inventories.TYPE_SUPPLY,
				// supplies work as tools
				AsTool: true,
			},
			UserId: userId,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuySuppliesRpcResponse{
		InventorySupplyKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
