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

	err = SetupSystemUsers(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = SetupSystemActivityExperiences(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupSystemRewards(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupCropRandomness(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupStarterConfigure(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupTokenConfigure(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupSpinConfigure(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupSupplies(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupDailyRewards(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = SetupSpins(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
