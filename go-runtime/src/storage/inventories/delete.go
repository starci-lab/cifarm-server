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
		err := nk.StorageDelete(ctx, []*runtime.StorageDelete{
			{
				Key:        params.Key,
				Collection: _constants.COLLECTION_INVENTORIES,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	}

	inventory.Quantity -= params.Quantity
	WriteInventoryObject(ctx, logger, db, nk, *inventory)
	return nil
}
