package farming_tools

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeFarmingTools(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexFarmingTools(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	return nil
}
