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
	TotalCost int64 `json:"totalCost"`
}

func BuySeedRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

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
		Metadata: map[string]interface{}{
			"name": "Buy seeds",
			"key":  params.Key,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = collections_inventories.Write(ctx,
		logger, db, nk,
		collections_inventories.WriteParams{
			Inventory: collections_inventories.Inventory{
				ReferenceId: params.Key,
				Quantity:    params.Quantity,
				Type:        collections_inventories.TYPE_SEED,
			},
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	_value, err := json.Marshal(BuySeedRpcResponse{
		TotalCost: totalCost,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
