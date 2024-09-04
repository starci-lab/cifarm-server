package seed_growth

import (
	_placed_items "cifarm-server/src/storage/placed_items"
	_system "cifarm-server/src/storage/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func HandleSeedGrowth(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	objects, err := _system.ReadSystemUsersObject(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	users, err := _system.ToSystemUsers(ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, userId := range users.UserIds {
		objects, err := _placed_items.ReadPlacedItemObjectsTypeFarmingTileIsPlanted(ctx, logger, db, nk, userId)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		farmingTiles, err := _placed_items.ToPlacedItems(ctx, logger, db, nk, objects)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		for _, farmingTile := range *farmingTiles {
			if farmingTile.SeedGrowthInfo.CurrentStage == farmingTile.SeedGrowthInfo.PlantSeed.GrowthStages {
				continue
			}
			time := int64(1)
			farmingTile.SeedGrowthInfo.TotalTimeElapsed += time
			farmingTile.SeedGrowthInfo.CurrentStageTimeElapsed += time

			if farmingTile.SeedGrowthInfo.CurrentStageTimeElapsed >= farmingTile.SeedGrowthInfo.PlantSeed.GrowthStageDuration {
				farmingTile.SeedGrowthInfo.CurrentStageTimeElapsed -= farmingTile.SeedGrowthInfo.PlantSeed.GrowthStageDuration
				farmingTile.SeedGrowthInfo.CurrentStage += 1
			}
		}
	}

	return err
}
