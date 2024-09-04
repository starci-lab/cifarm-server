package rpcs

import (
	_seed_growth "cifarm-server/src/crons/seed_growth"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeCron(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _seed_growth.RunSeedGrowthCron(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
