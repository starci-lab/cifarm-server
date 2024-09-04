package config

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteConfigPlayerMetdataObjectParams struct {
	PlayerMetadata _collections.PlayerMetadata
}

func WriteConfigPlayerMetdataObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteConfigPlayerMetdataObjectParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	value, err := json.Marshal(params.PlayerMetadata)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	write := &runtime.StorageWrite{
		UserID:          userId,
		Key:             _constants.KEY_PLAYER_METADATA,
		Collection:      _constants.COLLECTION_CONFIG,
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
