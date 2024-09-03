package rpcs

import (
	_daily_reward "cifarm-server/src/rpcs/daily_reward"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeRpcs(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := _daily_reward.InitializeStorageIndexLatestDailyRewardObject(ctx, logger, db, nk, initializer)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("go_healthcheck", HealthcheckRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("go_request_message", RequestMessageRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("go_daily_reward", _daily_reward.ClaimDailyRewardRpc)
	if err != nil {
		return err
	}

	return nil
}
