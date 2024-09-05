package collections_inventories

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
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
	inventory, err := collections_common.ToValue[Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if inventory == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
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
	err = Write(ctx, logger, db, nk, WriteParams{
		Inventory: *inventory,
		UserId:    params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
