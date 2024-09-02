package entities

import (
	_constants "cifarm-server/src/constants"
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

	plants := _collections.Plants{
		Items: []_collections.Plant{
			{
				Id:                  1,
				SeedPrice:           50,
				Name:                "Carrot",
				GrowthStageDuration: 1000 * 60 * 60, //1 hours
				GrowthStages:        5,
				Premium:             false,
				Perennial:           false,
				MinHarvestQuantity:  14,
				MaxHarvestQuantity:  20,
			},
			{
				Id:                          2,
				SeedPrice:                   100,
				Name:                        "Potato",
				GrowthStageDuration:         1000 * 60 * 60 * 2.5, //2.5 hours
				GrowthStages:                5,
				Premium:                     false,
				Perennial:                   false,
				MinHarvestQuantity:          16,
				MaxHarvestQuantity:          23,
				NextGrowthStageAfterHarvest: 1,
			},
		},
	}
	_plants, err := json.Marshal(plants)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      _constants.COLLECTION_ENTITIES,
			Key:             _constants.KEY_PLANTS,
			Value:           string(_plants),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
