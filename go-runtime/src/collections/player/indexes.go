package collections_player

import (
	collections_common "cifarm-server/src/collections/common"
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
	maxEntries := collections_common.MAX_ENTRIES
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func RegisterMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_METADATA
	collection := COLLECTION_NAME
	key := KEY_METADATA
	fields := []string{
		"chainKey",
		"accountAddress",
		"network",
	}
	sortableFields := []string{}
	maxEntries := collections_common.MAX_ENTRIES
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func RegisterMetadatas(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_METADATAS
	collection := COLLECTION_NAME
	key := KEY_METADATA
	fields := []string{
		"telegramData",
	}
	sortableFields := []string{}
	maxEntries := collections_common.MAX_ENTRIES
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
