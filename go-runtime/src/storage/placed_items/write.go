package placed_items

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

type WritePlacedItemObjectParams struct {
	PlacedItem _collections.PlacedItem
}

func WritePlacedItemObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WritePlacedItemObjectParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	value, err := json.Marshal(params.PlacedItem)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	write := &runtime.StorageWrite{
		UserID:          userId,
		Key:             uuid.NewString(),
		Collection:      _constants.COLLECTION_PLACED_ITEMS,
		Value:           string(value),
		PermissionRead:  2,
		PermissionWrite: 0,
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		write,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type WritePlacedItemObjectsParams struct {
	PlacedItems []_collections.PlacedItem
}

func WritePlacedItemObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WritePlacedItemObjectsParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	var writes []*runtime.StorageWrite

	for _, placedItem := range params.PlacedItems {
		value, err := json.Marshal(placedItem)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		write := &runtime.StorageWrite{
			UserID:          userId,
			Key:             uuid.NewString(),
			Collection:      _constants.COLLECTION_PLACED_ITEMS,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
