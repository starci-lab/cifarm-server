package inventories

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type DeleteInventoryObjectParams struct {
	Key      string `json:"key"`
	Quantity int    `json:"quantity"`
}

func DeleteInventoryObject(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteInventoryObjectParams,
) error {
	object, err := ReadInventoryObjectByKey(ctx, logger, db, nk, params.Key)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	inventory, err := ToInventory(ctx, logger, db, nk, object)
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
		logger.Info("equal")
		logger.Info(params.Key)
		err := nk.StorageDelete(ctx, []*runtime.StorageDelete{
			{
				UserID:     object.UserId,
				Key:        object.Key,
				Collection: _constants.COLLECTION_INVENTORIES,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	}
	logger.Info("nalz")
	inventory.Quantity -= params.Quantity
	err = WriteInventoryObject(ctx, logger, db, nk, *inventory)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
