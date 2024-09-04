package config

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadSystemUsersObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_SYSTEM_USERS
	query := ""
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		errMsg := "system users not found"
		logger.Error(errMsg)
		return nil, err
	}

	object := objects.Objects[0]
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
