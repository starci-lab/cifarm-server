package rpcs_users

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
	Accounts []Account `json:"accounts"`
}

type Account struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

func SearchUsersByValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params SearchUsersByValueParams,
) (*SearchUsersByValueResult, error) {
	query := `SELECT id, username FROM users WHERE username ILIKE concat('%', $1::TEXT, '%')
			  EXCEPT
              SELECT id, username FROM users WHERE id IN ($2);`
	rows, err := db.Query(query, params.Value, params.UserId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.UserId, &account.Username); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return &SearchUsersByValueResult{
		Accounts: accounts,
	}, nil
}

type SearchUsersRpcParams struct {
	Value string `json:"value"`
}

type SearchUsersRpcResponse struct {
	Accounts []Account `json:"accounts"`
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

	value, err := json.Marshal(SearchUsersRpcResponse{Accounts: result.Accounts})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
