package rpcs_shop

import (
	collections_animals "cifarm-server/src/collections/animals"
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuyAnimalRpcParams struct {
	Key string `json:"key"`
}

type BuyAnimalRpcResponse struct {
	Cost int64 `json:"cost"`
}

func BuyAnimalRpc(ctx context.Context,
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

	var params *BuyAnimalRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_animals.ReadByKey(ctx, logger, db, nk, collections_animals.ReadByKeyParams{
		Key: params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	animal, err := collections_common.ToValue[collections_animals.Animal](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if animal == nil {
		errMsg := "animal not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if animal.Premium {
		errMsg := "cannot buy premium animal"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		Amount: -animal.OffspringPrice,
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

	err = collections_inventories.Write(ctx,
		logger, db, nk,
		collections_inventories.WriteParams{
			Inventory: collections_inventories.Inventory{
				ReferenceId: params.Key,
				Quantity:    1,
				Type:        collections_inventories.TYPE_ANIMAL,
			},
			UserId: userId,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuyAnimalRpcResponse{
		Cost: animal.OffspringPrice,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
