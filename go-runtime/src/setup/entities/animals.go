package entities

import (
	_constants "cifarm-server/src/constants"
	_animals "cifarm-server/src/storage/animals"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

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
			Id:             _constants.ANIMAL_CHICKEN,
			Premium:        false,
			GrowthTime:     1000 * 60 * 60 * 7, //7 days
			YieldTime:      1000 * 60 * 60,     //1 days
		},
		{
			Id:         _constants.ANIMAL_COW,
			Premium:    true,
			GrowthTime: 1000 * 60 * 60 * 14, //14 days
			YieldTime:  1000 * 60 * 60 * 2,  //2 days
		},
	}

	err := _animals.WriteAnimalsObjects(ctx, logger, db, nk, _animals.WriteAnimalsObjectsParams{
		Animals: animals,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
