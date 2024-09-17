package crons_last_server_uptime

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Process(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	err := collections_system.WriteLastServerUptime(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
