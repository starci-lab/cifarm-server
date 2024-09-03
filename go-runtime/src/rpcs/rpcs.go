package rpcs

import (
	_wallets "cifarm-server/src/rpcs/wallets"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeRpcs(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("go_healthcheck", HealthcheckRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("go_request_message", RequestMessageRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("go_daily_reward", _wallets.ClaimDailyRewardRpc)
	if err != nil {
		return err
	}

	return nil
}
