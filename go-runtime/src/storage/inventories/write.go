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

type WriteInventoryObjectParams struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
	Type     int    `json:"int"`
}

func WriteInventoryObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteInventoryObjectParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	inventory, err := ReadInventoryObjectValue(ctx, logger, db, nk, ReadInventoryObjectParams{
		Id: params.Id,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if inventory != nil {
		inventory.Quantity += params.Quantity
	} else {
		inventory = &_collections.Inventory{
			Id:       params.Id,
			Type:     params.Type,
			Quantity: params.Quantity,
		}
	}

	_inventory, err := json.Marshal(inventory)
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
	return nil
}
