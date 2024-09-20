package rpcs_miscellaneous

import (
	collections_common "cifarm-server/src/collections/common"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

func ForceCentralBroadcastInstantlyRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	object, err := collections_system.ReadMatchInfo(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	matchInfo, err := collections_common.ToValue[collections_system.MatchInfo](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	_, err = nk.MatchSignal(ctx, matchInfo.CentralMatchId, userId)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
