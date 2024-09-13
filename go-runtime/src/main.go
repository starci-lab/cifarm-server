package main

import (
	"cifarm-server/src/auth"
	storage "cifarm-server/src/collections"
	"cifarm-server/src/crons"
	"cifarm-server/src/matches"
	"cifarm-server/src/rpcs"
	"cifarm-server/src/setup"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := storage.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = auth.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = rpcs.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = matches.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = setup.Initialize(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = crons.Initialize(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
