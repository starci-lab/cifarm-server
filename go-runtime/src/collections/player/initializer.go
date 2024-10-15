package collections_player

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := RegisterUserId(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = RegisterMetadata(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = RegisterMetadatas(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
