package farming_tiles

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeFarmingTiles(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexFarmingTiles(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
