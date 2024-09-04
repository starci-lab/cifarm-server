package rpcs

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func HealthcheckRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	return "ok", nil
}
