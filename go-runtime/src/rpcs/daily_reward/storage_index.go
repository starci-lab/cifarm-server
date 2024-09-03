package daily_reward

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageIndexLatestDailyRewardObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := _constants.STORAGE_INDEX_LATEST_DAILY_REWARD_OBJECTS
	collection := _constants.COLLECTION_REWARDS
	key := ""
	fields := []string{
		"days",
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
