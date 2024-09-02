package main

import (
	auth "cifarm-server/src/auth"
	rpcs "cifarm-server/src/rpcs"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Hello World!")
	err := initializer.RegisterBeforeAuthenticateCustom(auth.BeforeAuthenticate)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("go_healthcheck", RpcHealthcheck)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("go_request_message", rpcs.RequestMessageRpc)
	if err != nil {
		return err
	}

	return nil
}
