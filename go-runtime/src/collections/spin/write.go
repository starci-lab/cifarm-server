package collections_spin

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteManyParams struct {
	Spins []Spin `json:"spins"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, spin := range params.Spins {
		key := spin.Key
		spin.Key = ""
		value, err := json.Marshal(spin)
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
