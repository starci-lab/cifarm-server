package rpcs_users

import (
	"cifarm-server/src/friends"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type CheckFriendByUserIdRpcParams struct {
	UserId string `json:"userId"`
}

type CheckFriendByUserIdRpcResponse struct {
	Result bool `json:"result"`
}

func CheckFriendByUserIdRpc(
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

	var params *CheckFriendByUserIdRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	result, err := friends.CheckFriendByUserId(ctx, logger, db, nk, friends.CheckFriendByUserIdParams{
		UserId:       userId,
		FriendUserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(CheckFriendByUserIdRpcResponse{Result: result})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
