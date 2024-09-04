package main

import (
	_auth "cifarm-server/src/auth"
	_crons "cifarm-server/src/crons"
	_rpcs "cifarm-server/src/rpcs"
	_setup "cifarm-server/src/setup"
	_storage "cifarm-server/src/storage"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _storage.InitializeStorage(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _setup.InitializeSetup(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _auth.InitializeAuth(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _rpcs.InitializeRpcs(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _crons.InitializeCrons(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
