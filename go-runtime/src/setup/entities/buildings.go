package setup_entities

import (
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

	buildings := []collections_buildings.Bulding{
		{
			Key:   collections_buildings.KEY_COOP,
			Price: 1000,
		},
		{
			Key:   collections_buildings.KEY_PASTURE,
			Price: 2500,
		},
		{
			Key: collections_buildings.KEY_HOME,
		},
	}

	err := collections_buildings.WriteMany(ctx, logger, db, nk, collections_buildings.WriteManyParams{
		Buldings: buildings,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
