package setup_entities

import (
	collections_supplies "cifarm-server/src/collections/supplies"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSupplies(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	supplies := []collections_supplies.Supply{
		{
			Key:             collections_supplies.KEY_BASIC_FERTILIZER,
			Price:           50,
			AvailableInShop: true,
			Type:            collections_supplies.TYPE_FERTILIZER,
			FertilizerEffect: collections_supplies.FertilizerEffect{
				TimeReduce: 60 * 30, // 30 minutes
			},
		},
		{
			Key:             collections_supplies.KEY_ANIMAL_FEED,
			Price:           50,
			AvailableInShop: true,
			Type:            collections_supplies.TYPE_ANIMAL_FEED,
		},
	}

	err := collections_supplies.WriteMany(ctx, logger, db, nk, collections_supplies.WriteManyParams{
		Supplies: supplies,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
