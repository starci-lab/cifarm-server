package setup_entities

import (
	collections_crops "cifarm-server/src/collections/crops"
	collections_spin "cifarm-server/src/collections/spin"
	collections_supplies "cifarm-server/src/collections/supplies"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSpins(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	spins := []collections_spin.Spin{
		//100 golds
		{
			Key:          "gold1",
			Type:         collections_spin.TYPE_GOLD,
			GoldAmount:   100,
			ThresholdMin: 0,
			ThresholdMax: 0.2,
		},
		//250 golds
		{
			Key:          "gold2",
			Type:         collections_spin.TYPE_GOLD,
			GoldAmount:   250,
			ThresholdMin: 0.2,
			ThresholdMax: 0.35,
		},
		//500 golds
		{
			Key:          "gold3",
			Type:         collections_spin.TYPE_GOLD,
			GoldAmount:   500,
			ThresholdMin: 0.35,
			ThresholdMax: 0.45,
		},
		//1000 golds
		{
			Key:          "gold4",
			Type:         collections_spin.TYPE_GOLD,
			GoldAmount:   200,
			ThresholdMin: 0.45,
			ThresholdMax: 0.5,
		},
		//2 pineapple seeds
		{
			Key:          collections_crops.KEY_PINEAPPLE,
			Type:         collections_spin.TYPE_SEED,
			Quantity:     2,
			ThresholdMin: 0.5,
			ThresholdMax: 0.65,
		},
		//2 watermelon seeds
		{
			Key:          collections_crops.KEY_WATERMELON,
			Type:         collections_spin.TYPE_SEED,
			Quantity:     2,
			ThresholdMin: 0.65,
			ThresholdMax: 0.8,
		},
		//4 fertilizers
		{
			Key:          collections_supplies.KEY_BASIC_FERTILIZER,
			Type:         collections_spin.TYPE_SUPPLY,
			Quantity:     4,
			ThresholdMin: 0.8,
			ThresholdMax: 0.99,
		},
		//15 $CARROT
		{
			Key:          "token",
			Type:         collections_spin.TYPE_TOKEN,
			TokenAmount:  15,
			ThresholdMin: 0.99,
			ThresholdMax: 1,
		},
	}

	err := collections_spin.WriteMany(ctx, logger, db, nk, collections_spin.WriteManyParams{
		Spins: spins,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
