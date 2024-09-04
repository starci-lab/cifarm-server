package placed_items

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializePlacedItems(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexPlacedItems(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
