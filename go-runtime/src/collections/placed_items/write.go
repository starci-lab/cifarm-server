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
}

type WriteResult struct {
	Key string `json:"key"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) (*WriteResult, error) {
	key := params.PlacedItem.Key
	if key == "" {
		key = uuid.NewString()
	}
	params.PlacedItem.Key = ""

	value, err := json.Marshal(params.PlacedItem)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	acks, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return &WriteResult{
		Key: acks[0].Key,
	}, nil
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
		key := uuid.NewString()
		placedItem.Key = key

		value, err := json.Marshal(placedItem)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		write := &runtime.StorageWrite{
			Collection:      COLLECTION_NAME,
			Key:             key,
			Value:           string(value),
			UserID:          params.UserId,
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
