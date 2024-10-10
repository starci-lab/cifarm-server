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
			Key:             collections_buildings.KEY_CHICKEN_COOP,
			AvailableInShop: true,
			AnimalKey:       collections_animals.KEY_CHICKEN,
			MaxUpgrade:      3,
			Price:           2000,
			Capacity:        4,
		},
		{
			Key:             collections_buildings.KEY_COW_PASTURE,
			AvailableInShop: true,
			MaxUpgrade:      3,
			AnimalKey:       collections_animals.KEY_COW,
			Price:           3000,
			Capacity:        4,
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
