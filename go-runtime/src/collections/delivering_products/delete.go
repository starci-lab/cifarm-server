package collections_delivering_products

import (
	collections_common "cifarm-server/src/collections/common"
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
		errMsg := "delivery product not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	deliveringProduct, err := collections_common.ToValue[DeliveringProduct](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if params.Quantity > deliveringProduct.Quantity {
		errMsg := "cannot delete more than the available quantity in delivering product"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	if params.Quantity == deliveringProduct.Quantity {
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
	deliveringProduct.Quantity -= params.Quantity
	value, err := json.Marshal(deliveringProduct)
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

type DeleteManyParams struct {
	Keys   []string `json:"keys"`
	UserId string   `json:"userId"`
}

func DeleteMany(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteManyParams,
) error {
	var deletes []*runtime.StorageDelete
	for _, key := range params.Keys {
		deletes = append(deletes, &runtime.StorageDelete{
			Collection: COLLECTION_NAME,
			Key:        key,
			UserID:     params.UserId,
		})
	}
	err := nk.StorageDelete(ctx, deletes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
