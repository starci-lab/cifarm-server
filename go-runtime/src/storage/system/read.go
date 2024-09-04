package system

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadSystemUsersObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Key:        _constants.KEY_USERS,
			Collection: _constants.COLLECTION_SYSTEM,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects) == 0 {
		errMsg := "system users not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	object := objects[0]
	return object, nil
}

func ToSystemUsers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.Users, error) {
	if object == nil {
		return nil, nil
	}
	var users *_collections.Users
	err := json.Unmarshal([]byte(object.Value), &users)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return users, nil
}
