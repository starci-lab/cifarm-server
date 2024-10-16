package collections_friend_requests

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Register(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
		"referenceKey",
		"type",
		"premium",
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

func RegisterByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_BY_TOKEN_ID
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
		"tokenId",
		"referenceKey",
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

func RegisterByUserId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_UNIQUE
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
		"referenceKey",
		"unique",
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

func RegisterAvailable(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_AVAILABLE
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
		"isPlaced",
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
