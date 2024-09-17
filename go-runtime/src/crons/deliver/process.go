package crons_deliver

import (
	collections_common "cifarm-server/src/collections/common"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	collections_market_pricings "cifarm-server/src/collections/market-pricings"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Process(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	logger.Info("Delivering...")
	object, err := collections_system.ReadUsers(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	users, err := collections_common.ToValue[collections_system.Users](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, userId := range users.UserIds {
		go func() error {
			objects, err := collections_delivering_products.ReadMany(ctx, logger, db, nk, collections_delivering_products.ReadManyParams{
				UserId: userId,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			delivery_products, err := collections_common.ToValues2[collections_delivering_products.DeliveringProduct](ctx, logger, db, nk, objects)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			var totalGoldAmount int64
			var totalTokenAmount float64

			//get the key, delete the deliverings, then add money
			for _, delivery_product := range delivery_products {
				//ref to the reference
				marketPricingObject, err := collections_market_pricings.ReadByKey(ctx, logger, db, nk, collections_market_pricings.ReadByKeyParams{
					Key: delivery_product.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				marketPricing, err := collections_common.ToValue[collections_market_pricings.MarketPricing](ctx, logger, db, nk, marketPricingObject)
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				if !delivery_product.IsPremium {
					totalGoldAmount += marketPricing.BasicAmount * int64(delivery_product.Quantity)
				} else {
					totalTokenAmount += marketPricing.PremiumAmount * float64(delivery_product.Quantity)
				}
			}

			//update wallet
			err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
				UserId: userId,
				Amount: totalGoldAmount,
				Metadata: map[string]interface{}{
					"name": "Basic Delivery",
					"time": time.Now().Format(time.RFC850),
				},
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			//the other one might call api to process mint, peripery might be a wait
			return nil
		}()
	}
	return nil
}
