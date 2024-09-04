package inventories

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

func WriteInventoryObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	inventory _collections.Inventory,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	object, err := ReadInventoryObject(ctx, logger, db, nk, ReadInventoryObjectParams{
		Id: inventory.Id,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if object != nil {
		inventory, err := ToInventory(ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		inventory.Quantity += inventory.Quantity
		_inventory, err := json.Marshal(inventory)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
			{
				Collection:      _constants.COLLECTION_INVENTORIES,
				Key:             object.Key,
				UserID:          userId,
				Value:           string(_inventory),
				PermissionRead:  1,
				PermissionWrite: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	} else {
		_inventory, err := json.Marshal(
			inventory,
		)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
			{
				Collection:      _constants.COLLECTION_INVENTORIES,
				Key:             uuid.NewString(),
				UserID:          userId,
				Value:           string(_inventory),
				PermissionRead:  1,
				PermissionWrite: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	return nil
}
