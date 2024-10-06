package rpcs_community

import (
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
) (*User, error) {
	query := `SELECT id, username FROM users WHERE id NOT IN ($1, '00000000-0000-0000-0000-000000000000'::UUID) ORDER BY RANDOM() LIMIT 1;`
	rows, err := db.Query(query, params.UserId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.Username); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}
	return &user, nil
}

type GetRandomUserRpcResponse struct {
	User User `json:"user"`
}

func GetRandomUserRpc(
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
	user, err := GetRandomUser(ctx, logger, db, nk, GetRandomUserParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(GetRandomUserRpcResponse{User: *user})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
