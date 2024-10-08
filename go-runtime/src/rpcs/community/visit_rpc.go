package rpcs_community

import (
	collections_config "cifarm-server/src/collections/config"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VisitRpcParams struct {
	UserId string `json:"userId"`
}

func VisitRpc(
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

	var params *VisitRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	visitState := collections_config.VisitState{
		UserId: params.UserId,
	}
	err = collections_config.WriteVisitState(ctx, logger, db, nk, collections_config.WriteVisitStateParams{
		VisitState: visitState,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
