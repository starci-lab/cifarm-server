package rpcs_users

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type GetRandomUserParams struct {
	UserId string `json:"userId"`
}

func GetRandomUser(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetRandomUserParams,
) (string, error) {
	query := `SELECT id FROM users WHERE id NOT IN ($1) ORDER BY RANDOM() LIMIT 1;`
	rows, err := db.Query(query, params.UserId)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	defer rows.Close()

	var userId string
	for rows.Next() {
		if err := rows.Scan(userId); err != nil {
			logger.Error(err.Error())
			return "", err
		}
	}
	return userId, nil
}

type VisitRandomUserRpcResponse struct {
	UserId string `json:"userId"`
}

func VisitRandomUserRpc(
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
	randomUserId, err := GetRandomUser(ctx, logger, db, nk, GetRandomUserParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_config.ReadVisitState(ctx, logger, db, nk, collections_config.ReadVisitStateParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	visitState, err := collections_common.ToValue[collections_config.VisitState](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	visitState.UserId = randomUserId
	err = collections_config.WriteVisitState(ctx, logger, db, nk, collections_config.WriteVisitStateParams{
		VisitState: *visitState,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(VisitRandomUserRpcResponse{UserId: randomUserId})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
