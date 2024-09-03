package storage_queries

import (
	_animals "cifarm-server/src/storage_queries/animals"
	_daily_rewards "cifarm-server/src/storage_queries/daily_rewards"
	_inventories "cifarm-server/src/storage_queries/inventories"
	_plant_seeds "cifarm-server/src/storage_queries/plant_seeds"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageQueries(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _animals.InitializeStorageIndexAnimalsObjects(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = _plant_seeds.InitializeStorageIndexPlantSeedObjects(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = _inventories.InitializeStorageIndexInventoryObjects(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = _daily_rewards.InitializeStorageIndexLatestDailyRewardObject(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
