package storage

import (
	_animals "cifarm-server/src/storage/animals"
	_daily_rewards "cifarm-server/src/storage/daily_rewards"
	_inventories "cifarm-server/src/storage/inventories"
	_plant_seeds "cifarm-server/src/storage/plant_seeds"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorage(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
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
