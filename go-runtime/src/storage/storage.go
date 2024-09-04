package storage

import (
	_animals "cifarm-server/src/storage/animals"
	_config "cifarm-server/src/storage/config"
	_daily_rewards "cifarm-server/src/storage/daily_rewards"
	_farming_tiles "cifarm-server/src/storage/farming_tiles"
	_farming_tools "cifarm-server/src/storage/farming_tools"
	_inventories "cifarm-server/src/storage/inventories"
	_placed_items "cifarm-server/src/storage/placed_items"
	_plant_seeds "cifarm-server/src/storage/plant_seeds"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorage(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _animals.InitializeAnimals(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _plant_seeds.InitializePlantSeeds(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _inventories.InitializeInventory(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _daily_rewards.InitializeDailyReward(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _farming_tools.InitializeFarmingTools(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _farming_tiles.InitializeFarmingTiles(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _config.InitializeConfig(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _placed_items.InitializePlacedItems(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
