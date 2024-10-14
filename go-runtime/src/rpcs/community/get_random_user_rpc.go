package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

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
	object, err := collections_config.ReadMetadata(ctx, logger, db, nk, collections_config.ReadMetadataParams{
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	metadata, err := collections_common.ToValue[collections_config.Metadata](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	query := `SELECT id, username FROM users INNER JOIN storage ON users.id = storage.user_id
WHERE id NOT IN ('00000000-0000-0000-0000-000000000000'::UUID) 
AND key = 'metadata'
AND value->'telegramData'->>'userId' != $1
ORDER BY RANDOM() LIMIT 1;
`

	telegramUserId := strconv.Itoa(metadata.TelegramData.UserId)
	rows, err := db.Query(query, telegramUserId)
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
