package daily_rewards

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeDailyReward(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("claim_daily_reward", ClaimDailyRewardRpc)
	if err != nil {
		return err
	}
	return nil
}
