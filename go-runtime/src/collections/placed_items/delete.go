package collections_placed_items

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type DeleteParams struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}

func Delete(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteParams,
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

type DeleteByInventoryKeyParams struct {
	InventoryKey string `json:"inventoryKey"`
	UserId       string `json:"userId"`
}

func DeleteByInventoryKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteByInventoryKeyParams,
) error {
	object, err := ReadByInventoryKey(ctx, logger, db, nk, ReadByInventoryKeyParams{
		InventoryKey: params.InventoryKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = nk.StorageDelete(ctx, []*runtime.StorageDelete{
		{
			Key:        object.Key,
			Collection: COLLECTION_NAME,
			UserID:     params.UserId,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
