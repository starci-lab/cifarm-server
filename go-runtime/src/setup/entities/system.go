package setup_entities

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSystem(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	object, err := collections_system.ReadByKey(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object != nil {
		return nil
	}

	users := collections_system.Users{
		UserIds: []string{},
	}

	err = collections_system.Write(ctx, logger, db, nk, collections_system.WriteParams{
		Users: users,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
