package config

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageIndexSystemUsers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := _constants.STORAGE_INDEX_SYSTEM_USERS
	collection := _constants.COLLECTION_SYSTEM
	key := _constants.KEY_PLAYER_METADATA
	fields := []string{
		"userIds",
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
