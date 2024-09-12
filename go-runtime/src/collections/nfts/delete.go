package collections_nfts

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
	var writes []*runtime.StorageDelete
	for _, key := range params.Keys {
		writes = append(writes, &runtime.StorageDelete{
			Collection: COLLECTION_NAME,
			Key:        key,
			UserID:     params.UserId,
		})
	}
	err := nk.StorageDelete(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
