package setup_entities

import (
	collections_animals "cifarm-server/src/collections/animals"
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

	animals := []collections_animals.Animal{
		{
			Key:             collections_animals.KEY_CHICKEN,
			OffspringPrice:  1000,
			Premium:         false,
			GrowthTime:      60 * 60 * 7, //7 days
			YieldTime:       60 * 60,     //1 days
			AvailableInShop: true,
			HungerTime:      60 * 12, //12 hours
		},
		{
			Key:             collections_animals.KEY_COW,
			Premium:         true,
			GrowthTime:      60 * 60 * 14, //14 days
			YieldTime:       60 * 60 * 2,  //2 days
			AvailableInShop: false,
			HungerTime:      60 * 12, //12 hours
		},
	}

	err := collections_animals.WriteMany(ctx, logger, db, nk, collections_animals.WriteManyParams{
		Animals: animals,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
