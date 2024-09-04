package setup

import (
	_entities "cifarm-server/src/setup/entities"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeSetup(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	err := _entities.SetupPlants(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = _entities.SetupAnimals(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = _entities.SetupFarmingTools(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = _entities.SetupFarmingTiles(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
