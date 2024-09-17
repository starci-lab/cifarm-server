package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	collections_inventories "cifarm-server/src/collections/inventories"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type RetainProductsRpcParams struct {
	DeliveringProducts []collections_delivering_products.DeliveringProduct `json:"deliveringProducts"`
}

type RetainProductsRpcResponse struct {
	InventoryKeys []string `json:"inventoryKeys"`
}

func RetainProductsRpc(
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

	var params *RetainProductsRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var keys []string
	for _, deliveringProduct := range params.DeliveringProducts {
		//query delivery product
		object, err := collections_delivering_products.ReadByKey(ctx, logger, db, nk, collections_delivering_products.ReadByKeyParams{
			UserId: userId,
			Key:    deliveringProduct.Key,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		query, err := collections_common.ToValue[collections_delivering_products.DeliveringProduct](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		//delete the previous
		err = collections_delivering_products.Delete(ctx, logger, db, nk, collections_delivering_products.DeleteParams{
			Key:      deliveringProduct.Key,
			Quantity: deliveringProduct.Quantity,
			UserId:   userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		//write new inventories
		var inventoryType int
		switch query.Type {
		case collections_delivering_products.TYPE_PLANT:
			inventoryType = collections_inventories.TYPE_HARVESTED_PLANT
		default:
		}

		result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
			Inventory: collections_inventories.Inventory{
				ReferenceKey: query.ReferenceKey,
				Quantity:     deliveringProduct.Quantity,
				Type:         inventoryType,
				IsPremium:    true,
			},
			UserId: userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		keys = append(keys, result.Key)
	}

	value, err := json.Marshal(RetainProductsRpcResponse{
		InventoryKeys: keys,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
