package entities

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupFarmingTools(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	animals := _collections.Animals{
		Items: []_collections.Animal{
			{
				Id:             1,
				OffspringPrice: 1000,
				Name:           "Chicken",
				Premium:        false,
				GrowthTime:     1000 * 60 * 60 * 7, //7 days
				YieldTime:      1000 * 60 * 60,     //1 days
			},
			{
				Id:         1,
				Name:       "Cow",
				Premium:    true,
				GrowthTime: 1000 * 60 * 60 * 14, //14 days
				YieldTime:  1000 * 60 * 60 * 2,  //2 days
			},
		},
	}

	_animals, err := json.Marshal(animals)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      _constants.COLLECTION_ENTITIES,
			Key:             _constants.KEY_ANIMALS,
			Value:           string(_animals),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
