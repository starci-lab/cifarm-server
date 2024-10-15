package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_player "cifarm-server/src/collections/player"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

func ReturnRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	object, err := collections_player.ReadVisitState(ctx, logger, db, nk, collections_player.ReadVisitStateParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	visitState, err := collections_common.ToValue[collections_player.VisitState](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if visitState.UserId == "" {
		errMsg := "not visit"
		logger.Error(errMsg)
		return "", err
	}
	visitState.UserId = ""

	err = collections_player.WriteVisitState(ctx, logger, db, nk, collections_player.WriteVisitStateParams{
		VisitState: *visitState,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
