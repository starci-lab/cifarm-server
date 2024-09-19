package setup_entities

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule,
) error {
	err := SetupCrops(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupAnimals(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupTools(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupTiles(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupBuildings(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupSystem(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupMarketPricings(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
