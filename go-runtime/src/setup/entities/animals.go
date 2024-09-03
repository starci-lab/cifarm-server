package entities

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupAnimals(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	animals := []_collections.Animal{
		{
			OffspringPrice: 1000,
			Id:             _collections.ANIMAL_CHICKEN,
			Premium:        false,
			GrowthTime:     1000 * 60 * 60 * 7, //7 days
			YieldTime:      1000 * 60 * 60,     //1 days
		},
		{
			Id:         _collections.ANIMAL_COW,
			Premium:    true,
			GrowthTime: 1000 * 60 * 60 * 14, //14 days
			YieldTime:  1000 * 60 * 60 * 2,  //2 days
		},
	}

	var writes []*runtime.StorageWrite
	for _, animal := range animals {
		value, err := json.Marshal(animal)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_ANIMALS,
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
