package daily_reward

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeDailyReward(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := InitializeStorageIndexLatestDailyRewardObject(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("go_claim_daily_reward", ClaimDailyRewardRpc)
	if err != nil {
		return err
	}

	return nil
}
