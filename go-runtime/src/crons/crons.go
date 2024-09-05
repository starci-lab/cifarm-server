package crons

import (
	"cifarm-server/src/crons/last_server_uptime"
	crons_seed_growth "cifarm-server/src/crons/seed_growth"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	err := last_server_uptime.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = crons_seed_growth.Run(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
