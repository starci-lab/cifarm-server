package rpcs_community

import (
	collections_player "cifarm-server/src/collections/player"
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

	visitState := collections_player.VisitState{
		UserId: params.UserId,
	}
	err = collections_player.WriteVisitState(ctx, logger, db, nk, collections_player.WriteVisitStateParams{
		VisitState: visitState,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
