package plant_seeds

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WritePlantSeedObjectsParams struct {
	PlantSeeds []_collections.PlantSeed
}

func WritePlantSeedObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WritePlantSeedObjectsParams,
) error {
	var writes []*runtime.StorageWrite
	for _, plantSeed := range params.PlantSeeds {
		value, err := json.Marshal(plantSeed)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_PLANT_SEEDS,
			Key:             plantSeed.Id,
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
