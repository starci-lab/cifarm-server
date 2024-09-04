package config

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeSystem(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexSystemUsers(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
