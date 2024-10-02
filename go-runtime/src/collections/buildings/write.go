package collections_buildings

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteManyParams struct {
	Buildings []Building `json:"buildings"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, Building := range params.Buildings {
		key := Building.Key
		Building.Key = ""

		value, err := json.Marshal(Building)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		write := &runtime.StorageWrite{
			Key:             key,
			Collection:      COLLECTION_NAME,
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
