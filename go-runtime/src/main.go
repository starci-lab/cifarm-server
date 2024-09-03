package main

import (
	_auth "cifarm-server/src/auth"
	_rpcs "cifarm-server/src/rpcs"
	_setup "cifarm-server/src/setup"
	_storage_queries "cifarm-server/src/storage_queries"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _storage_queries.InitializeStorageQueries(ctx, logger, db, nk, initializer)
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
		return err
	}
	return nil
}
