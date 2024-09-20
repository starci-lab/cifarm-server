package setup_entities

import (
	collections_crops "cifarm-server/src/collections/crops"
	collections_market_pricings "cifarm-server/src/collections/market_pricings"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupMarketPricings(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	marketPricings := []collections_market_pricings.MarketPricing{
		{
			Key:           collections_crops.KEY_CARROT,
			BasicAmount:   4,
			PremiumAmount: 0.04,
		},
		{
			Key:           collections_crops.KEY_POTATO,
			BasicAmount:   8,
			PremiumAmount: 0.08,
		},
	}

	err := collections_market_pricings.WriteMany(ctx, logger, db, nk, collections_market_pricings.WriteManyParams{
		MarketPricings: marketPricings,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
