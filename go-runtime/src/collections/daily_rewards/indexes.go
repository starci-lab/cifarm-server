package collections_daily_rewards

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func RegisterHighestPossibleDay(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_HIGHEST_POSSIBLE_DAY
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
		"day",
	}
	sortableFields := []string{
		"day",
	}
	maxEntries := collections_common.MAX_ENTRIES
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
