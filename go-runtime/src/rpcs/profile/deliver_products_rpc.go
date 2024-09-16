package rpcs_profiles

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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
			return false, nil
		}
	}
	return true, nil
}

type DeliverProductsRpcParams struct {
	Inventories []collections_inventories.Inventory `json:"inventories"`
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

	//update the state of delivery
	object, err := collections_config.ReadDeliveryState(ctx, logger, db, nk, collections_config.ReadDeliveryStateParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object != nil {
		deliveryState, err := collections_common.ToValue[collections_config.DeliveryState](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		if deliveryState.Delivering {
			errMsg := "products delivering"
			logger.Error(errMsg)
			return "", errors.New(errMsg)
		}
	}
	deliveryState := collections_config.DeliveryState{
		Delivering: true,
	}

	err = collections_config.WriteDeliveryState(ctx, logger, db, nk, collections_config.WriteDeliveryStateParams{
		UserId:        userId,
		DeliveryState: deliveryState,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	for _, inventory := range params.Inventories {
		//process writes
		inventory.IsDelivering = true
		_, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
			Inventory: inventory,
			UserId:    userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
	}

	return "", nil
}
