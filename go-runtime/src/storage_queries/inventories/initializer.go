package inventories

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeInventory(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexInventoryObjects(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
