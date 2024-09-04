package system

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func WriteSystemUsersObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	users _collections.Users,
) error {
	value, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	write := &runtime.StorageWrite{
		Key:             _constants.KEY_USERS,
		Collection:      _constants.COLLECTION_SYSTEM,
		Value:           string(value),
		PermissionRead:  2,
		PermissionWrite: 0,
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		write,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
