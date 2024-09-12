package collections_daily_rewards

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func RegisterLatest(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := STORAGE_INDEX_LATEST
	collection := COLLECTION_NAME
	key := ""
	fields := []string{
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
