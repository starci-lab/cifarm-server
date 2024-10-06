package collections_config

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func RegisterUserId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_USER_ID
	collection := COLLECTION_NAME
	key := KEY_METADATA
	fields := []string{
		"accountAddress",
		"chainKey",
		"network",
	}
	sortableFields := []string{}
	maxEntries := 1
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
