package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	collections_inventories "cifarm-server/src/collections/inventories"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

type EnsureParams struct {
	UserId      string                              `json:"userId"`
	Inventories []collections_inventories.Inventory `json:"inventories"`
}

func Ensure(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params EnsureParams,
) (bool, error) {
	var keys []string
	for _, inventory := range params.Inventories {
		keys = append(keys, inventory.Key)
	}
	objects, err := collections_inventories.ReadMany(ctx, logger, db, nk, collections_inventories.ReadManyParams{
		UserId: params.UserId,
		Keys:   keys,
	})
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	queriedInventories, err := collections_common.ToValues2[collections_inventories.Inventory](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	for index, queriedInventory := range queriedInventories {
		// nếu số lượng trong cơ sở dữ liệu bé hơn
		if queriedInventory.Quantity < params.Inventories[index].Quantity {
			errMsg := fmt.Sprintf("quantity not enough: %s", queriedInventory.Key)
			logger.Error(errMsg)
			return false, nil
		}
		//nếu mà nó không deliveriable được
		if !queriedInventory.Deliverable {
			errMsg := fmt.Sprintf("not deliverialbe: %s", queriedInventory.Key)
			logger.Error(errMsg)
			return false, nil
		}
	}
	return true, nil
}

type DeliverProductsRpcParams struct {
	Inventories []collections_inventories.Inventory `json:"inventories"`
}

type DeliverProductsRpcResponse struct {
	DeliveryProductKeys []string `json:"deliveryProductKeys"`
}

func DeliverProductsRpc(
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

	var params *DeliverProductsRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//ensure enough item to deliver
	ensure, err := Ensure(ctx, logger, db, nk, EnsureParams{
		UserId:      userId,
		Inventories: params.Inventories,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if !ensure {
		errMsg := "not enough quantity to deliver"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var keys []string
	for _, inventory := range params.Inventories {
		//query again to track data
		object, err := collections_inventories.ReadByKey(ctx, logger, db, nk, collections_inventories.ReadByKeyParams{
			UserId: userId,
			Key:    inventory.Key,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		query, err := collections_common.ToValue[collections_inventories.Inventory](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		//delete the previous
		err = collections_inventories.Delete(ctx, logger, db, nk, collections_inventories.DeleteParams{
			Key:      inventory.Key,
			Quantity: inventory.Quantity,
			UserId:   userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		//write new delivering products
		var productType int
		switch query.Type {
		case collections_inventories.TYPE_HARVESTED_PLANT:
			productType = collections_delivering_products.TYPE_PLANT
		default:
		}

		result, err := collections_delivering_products.Write(ctx, logger, db, nk, collections_delivering_products.WriteParams{
			DeliveringProduct: collections_delivering_products.DeliveringProduct{
				ReferenceKey: query.ReferenceKey,
				Quantity:     inventory.Quantity,
				Type:         productType,
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

	value, err := json.Marshal(DeliverProductsRpcResponse{
		DeliveryProductKeys: keys,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
