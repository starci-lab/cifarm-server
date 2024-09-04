package animals

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func WriteAnimalsObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	animals []_collections.Animal,
) error {
	var writes []*runtime.StorageWrite
	for _, animal := range animals {
		value, err := json.Marshal(animal)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      _constants.COLLECTION_ANIMALS,
			Key:             animal.Id,
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
