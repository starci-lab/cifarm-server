package hooks

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

func AfterDeleteAccount(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	err := collections_system.DeleteUser(ctx, logger, db, nk, collections_system.DeleteUserParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
