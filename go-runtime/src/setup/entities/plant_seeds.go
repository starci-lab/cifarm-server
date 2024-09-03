package entities

import (
	_plant_seeds "cifarm-server/src/storage/plant_seeds"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupPlants(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	plantSeeds := []_collections.PlantSeed{
		{
			Id:                  _collections.PLANT_SEED_CARROT_SEED,
			SeedPrice:           50,
			GrowthStageDuration: 1000 * 60 * 60, //1 hours
			GrowthStages:        5,
			Premium:             false,
			Perennial:           false,
			MinHarvestQuantity:  14,
			MaxHarvestQuantity:  20,
		},
		{
			Id:                          _collections.PLANT_SEED_POTATO_SEED,
			SeedPrice:                   100,
			GrowthStageDuration:         1000 * 60 * 60 * 2.5, //2.5 hours
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
		},
	}

	err := _plant_seeds.WritePlantSeedObjects(ctx, logger, db, nk, _plant_seeds.WritePlantSeedObjectsParams{
		PlantSeeds: plantSeeds,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
