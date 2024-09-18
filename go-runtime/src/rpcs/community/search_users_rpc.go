package rpcs_community

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type SearchUsersByValueParams struct {
	UserId string `json:"userId"`
	Value  string `json:"value"`
}

type SearchUsersByValueResult struct {
	Users []User `json:"users"`
}

func SearchUsersByValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params SearchUsersByValueParams,
) (*SearchUsersByValueResult, error) {
	query := `SELECT id, username FROM users WHERE username ILIKE concat('%', $1::TEXT, '%') AND id NOT IN ($2, '00000000-0000-0000-0000-000000000000'::UUID);`
	rows, err := db.Query(query, params.Value, params.UserId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserId, &user.Username); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return &SearchUsersByValueResult{
		Users: users,
	}, nil
}

type SearchUsersRpcParams struct {
	Value string `json:"value"`
}

type SearchUsersRpcResponse struct {
	Users []User `json:"users"`
}

func SearchUserRpc(
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

	var params *SearchUsersRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if params.Value == "" {
		errMsg := "search value cannot be empty"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	result, err := SearchUsersByValue(ctx, logger, db, nk, SearchUsersByValueParams{
		UserId: userId,
		Value:  params.Value,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(SearchUsersRpcResponse{Users: result.Users})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
