package main

import (
	_avalanche "cifarm-server/src/auth/avalanche"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Hello World!")
	err := initializer.RegisterBeforeAuthenticateCustom(_avalanche.BeforeAvalancheAuth)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("go_healthcheck", RpcHealthcheck)
	if err != nil {
		return err
	}

	return nil
}
