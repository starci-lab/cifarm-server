package rpcs_community

import (
	"cifarm-server/src/friends"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type CheckFriendByUsernameRpcParams struct {
	Username string `json:"username"`
}

type CheckFriendByUsernameRpcResponse struct {
	Result bool `json:"result"`
}

func CheckFriendByUsernameRpc(
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

	var params *CheckFriendByUsernameRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	result, err := friends.CheckFriendByUsername(ctx, logger, db, nk, friends.CheckFriendByUsernameParams{
		UserId:         userId,
		FriendUsername: params.Username,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(CheckFriendByUsernameRpcResponse{Result: result})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
