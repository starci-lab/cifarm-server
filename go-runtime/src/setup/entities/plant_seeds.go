package entities

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

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
			Id:                  "carrot",
			SeedPrice:           50,
			GrowthStageDuration: 1000 * 60 * 60, //1 hours
			GrowthStages:        5,
			Premium:             false,
			Perennial:           false,
			MinHarvestQuantity:  14,
			MaxHarvestQuantity:  20,
		},
		{
			Id:                          "potato",
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

	var writes []*runtime.StorageWrite
	for _, plantSeed := range plantSeeds {
		value, err := json.Marshal(plantSeed)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_PLANT_SEEDS,
			Key:             plantSeed.Id,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
