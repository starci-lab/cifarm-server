package main

import (
	auth "cifarm-server/src/auth"
	rpcs "cifarm-server/src/rpcs"
	setup "cifarm-server/src/setup"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	err := setup.Setup(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = initializer.RegisterBeforeAuthenticateCustom(auth.BeforeAuthenticate)
	if err != nil {
		return err
	}

	err = initializer.RegisterAfterAuthenticateCustom(auth.AfterAuthenticate)
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
