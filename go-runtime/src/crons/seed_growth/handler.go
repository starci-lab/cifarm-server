package seed_growth

import (
	_placed_items "cifarm-server/src/storage/placed_items"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HandleSeedGrowthParams struct {
	UserId string `json:"userId"`
	Dumb   int    `json:"-"`
}

func HandleSeedGrowth(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params HandleSeedGrowthParams,
) error {
	objects, err := _placed_items.ReadOwnedPlacedItemObjectsTypeFarmingTileIsPlanted(ctx, logger, db, nk, params.UserId)
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

	return err
}
