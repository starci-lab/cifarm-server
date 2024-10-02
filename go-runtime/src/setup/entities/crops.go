package setup_entities

import (
	collections_crops "cifarm-server/src/collections/crops"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupCrops(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	crops := []collections_crops.Crop{
		{
			Key:                       collections_crops.KEY_CARROT,
			Price:                     50,
			GrowthStageDuration:       60 * 60, //1 hours
			GrowthStages:              5,
			Premium:                   false,
			Perennial:                 false,
			MinHarvestQuantity:        14,
			MaxHarvestQuantity:        20,
			BasicHarvestExperiences:   12,
			PremiumHarvestExperiences: 60,
			AvailableInShop:           true,
		},
		{
			Key:                         collections_crops.KEY_POTATO,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
			AvailableInShop:             true,
		},
		{
			Key:                         collections_crops.KEY_CUCUMBER,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
			AvailableInShop:             true,
		},
		{
			Key:                         collections_crops.KEY_PINEAPPLE,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
			AvailableInShop:             true,
		},
		{
			Key:                         collections_crops.KEY_WATERMELON,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
			AvailableInShop:             true,
		},
		{
			Key:                         collections_crops.KEY_PINEAPPLE,
			Price:                       100,
			GrowthStageDuration:         60 * 60 * 2.5, //2.5 hours  60 * 60 * 2.5
			GrowthStages:                5,
			Premium:                     false,
			Perennial:                   false,
			MinHarvestQuantity:          16,
			MaxHarvestQuantity:          23,
			NextGrowthStageAfterHarvest: 1,
			BasicHarvestExperiences:     21,
			PremiumHarvestExperiences:   110,
			AvailableInShop:             true,
		},
	}

	err := collections_crops.WriteMany(ctx, logger, db, nk, collections_crops.WriteManyParams{
		Crops: crops,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
