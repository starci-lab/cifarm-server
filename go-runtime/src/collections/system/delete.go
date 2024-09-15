package collections_system

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type DeleteUserParams struct {
	UserId string `json:"userId"`
}

func DeleteUser(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteUserParams,
) error {
	object, err := ReadUsers(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	users, err := collections_common.ToValue[Users](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var updatedUserIds []string
	for _, userId := range users.UserIds {
		if userId != params.UserId {
			updatedUserIds = append(updatedUserIds, userId)
		}
	}

	users.UserIds = updatedUserIds
	value, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_USERS,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
