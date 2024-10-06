package rpcs_internal

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func CheckPermission(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) bool {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return false
	}
	account, err := nk.AccountGetId(ctx, userId)
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	authenticationId, err := config.AuthenticationId(ctx, logger, db, nk)
	if err != nil {
		return false
	}
	if account.Devices[0].Id != authenticationId {
		errMsg := "permission denied"
		logger.Error(errMsg)
		return false
	}
	return true
}
