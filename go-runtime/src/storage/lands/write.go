package lands

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteLandObjectsParams struct {
	Lands []_collections.Land
}

func WriteLandObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteLandObjectsParams,
) error {
	var writes []*runtime.StorageWrite
	for _, land := range params.Lands {
		value, err := json.Marshal(land)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_LANDS,
			Key:             land.Id,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
