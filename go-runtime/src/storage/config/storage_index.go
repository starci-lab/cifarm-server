package config

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageIndexConfigPlayerMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := _constants.STORAGE_INDEX_CONFIG_PLAYER_METADATA
	collection := _constants.COLLECTION_CONFIG
	key := _constants.KEY_PLAYER_METADATA
	fields := []string{
		"chain",
		"address",
	}
	sortableFields := []string{}
	maxEntries := 100
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
