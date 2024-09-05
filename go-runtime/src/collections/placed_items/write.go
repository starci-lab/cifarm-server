package collections_placed_items

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	PlacedItem PlacedItem `json:"placedItem"`
	UserId     string     `json:"userId"`
	Key        string     `json:"key"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) error {
	value, err := json.Marshal(params.PlacedItem)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if params.Key == "" {
		params.Key = uuid.NewString()
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             params.Key,
			UserID:          params.UserId,
			Value:           string(value),
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

type WriteManyParams struct {
	PlacedItems []PlacedItem `json:"placedItems"`
	UserId      string       `json:"userId"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, placedItem := range params.PlacedItems {
		value, err := json.Marshal(placedItem)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      COLLECTION_NAME,
			Key:             uuid.NewString(),
			Value:           string(value),
			UserID:          params.UserId,
			PermissionRead:  1,
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
