package collections_friend_requests

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type DeleteParams struct {
	Key      string `json:"key"`
	Quantity int    `json:"quantity"`
	UserId   string `json:"userId"`
}

func Delete(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteParams,
) error {
	object, err := ReadByKey(ctx, logger, db, nk, ReadByKeyParams{
		Key:    params.Key,
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if object == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	inventory, err := collections_common.ToValue[Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if params.Quantity > inventory.Quantity {
		errMsg := "cannot delete more than the available quantity in inventory"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	if params.Quantity == inventory.Quantity {
		err := nk.StorageDelete(ctx, []*runtime.StorageDelete{
			{
				Collection: COLLECTION_NAME,
				Key:        object.Key,
				UserID:     object.UserId,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	}
	inventory.Quantity -= params.Quantity
	value, err := json.Marshal(inventory)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             object.Key,
			UserID:          object.UserId,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type DeleteUniqueParams struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}

func DeleteUnique(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteUniqueParams,
) error {
	err := nk.StorageDelete(ctx, []*runtime.StorageDelete{
		{
			Collection: COLLECTION_NAME,
			Key:        params.Key,
			UserID:     params.UserId,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_placed_items.DeleteByInventoryKey(ctx, logger, db, nk, collections_placed_items.DeleteByInventoryKeyParams{
		InventoryKey: params.Key,
		UserId:       params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
