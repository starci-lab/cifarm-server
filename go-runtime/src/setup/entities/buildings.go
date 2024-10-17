package setup_entities

import (
	collections_animals "cifarm-server/src/collections/animals"
	collections_buildings "cifarm-server/src/collections/buildings"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupBuildings(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	buildings := []collections_buildings.Building{
		{
			Key:             collections_buildings.KEY_COOP,
			AvailableInShop: true,
			Type:            collections_animals.TYPE_POULTRY,
			MaxUpgrade:      2,
			Price:           2000,
			Upgrades: map[int]collections_buildings.Upgrade{
				0: {
					UpgradePrice: 0,
					Capacity:     3,
				},
				1: {
					UpgradePrice: 1000,
					Capacity:     5,
				},
				2: {
					UpgradePrice: 2000,
					Capacity:     10,
				},
			},
		},
		{
			Key:             collections_buildings.KEY_PASTURE,
			AvailableInShop: true,
			MaxUpgrade:      2,
			Type:            collections_animals.TYPE_LIVESTOCK,
			Price:           3000,
			Upgrades: map[int]collections_buildings.Upgrade{
				0: {
					UpgradePrice: 0,
					Capacity:     3,
				},
				1: {
					UpgradePrice: 1000,
					Capacity:     5,
				},
				2: {
					UpgradePrice: 2000,
					Capacity:     10,
				},
			},
		},
		{
			Key: collections_buildings.KEY_HOME,
		},
	}

	err := collections_buildings.WriteMany(ctx, logger, db, nk, collections_buildings.WriteManyParams{
		Buildings: buildings,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
