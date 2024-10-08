package crons

import (
	crons_animal_produce "cifarm-server/src/crons/animal_produce"
	crons_deliver "cifarm-server/src/crons/deliver"
	crons_energy_gain "cifarm-server/src/crons/energy_gain"
	crons_last_server_uptime "cifarm-server/src/crons/last_server_uptime"
	crons_seed_growth "cifarm-server/src/crons/seed_growth"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	err := crons_last_server_uptime.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = crons_seed_growth.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = crons_deliver.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = crons_energy_gain.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = crons_animal_produce.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
