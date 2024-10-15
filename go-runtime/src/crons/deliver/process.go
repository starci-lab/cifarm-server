package crons_deliver

import (
	collections_common "cifarm-server/src/collections/common"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	collections_market_pricings "cifarm-server/src/collections/market_pricings"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"sync"
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

	var wg sync.WaitGroup
	for _, userId := range users.UserIds {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			objects, err := collections_delivering_products.ReadMany(ctx, logger, db, nk, collections_delivering_products.ReadManyParams{
				UserId: userId,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			deliveryProducts, err := collections_common.ToValues2[collections_delivering_products.DeliveringProduct](ctx, logger, db, nk, objects)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			//delete all delivers
			var keys []string
			var totalGoldAmount int64
			var totalTokenAmount float64

			//get the key, delete the deliverings, then add money

			for _, deliveryProduct := range deliveryProducts {
				keys = append(keys, deliveryProduct.Key)

				//ref to the reference
				marketPricingObject, err := collections_market_pricings.ReadByKey(ctx, logger, db, nk, collections_market_pricings.ReadByKeyParams{
					Key: deliveryProduct.ReferenceKey,
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
				totalGoldAmount += marketPricing.BasicAmount * int64(deliveryProduct.Quantity)
				if deliveryProduct.Premium {
					totalTokenAmount += marketPricing.PremiumAmount * float64(deliveryProduct.Quantity)
				}
			}

			//delete
			err = collections_delivering_products.DeleteMany(ctx, logger, db, nk, collections_delivering_products.DeleteManyParams{
				UserId: userId,
				Keys:   keys,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			//update wallet
			err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
				UserId:      userId,
				GoldAmount:  totalGoldAmount,
				TokenAmount: totalTokenAmount,
				Metadata: map[string]interface{}{
					"name": "Basic Delivery",
					"time": time.Now().Format(time.RFC850),
				},
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			// //get the metadata
			// metadataObject, err := collections_player.ReadMetadata(ctx, logger, db, nk, collections_player.ReadMetadataParams{
			// 	UserId: userId,
			// })
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	return err
			// }
			// metadata, err := collections_common.ToValue[collections_player.Metadata](ctx, logger, db, nk, metadataObject)
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	return err
			// }

			// //the other one might call api to process mint, peripery might be a wait
			// minterPrivatekey, err := config.MinterPrivateKey(ctx, logger, db, nk)
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	return err
			// }
			// utilityTokenAddress, err := config.UtilityTokenAddress(ctx, logger, db, nk)
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	return err
			// }

			// //
			// //GasCheck is required (future plan)
			// //ect
			// //maybe do later with response, such as notifcation,...
			// _, err = services_periphery_api_token.Mint(ctx, logger, db, nk, &services_periphery_api_token.MintRequestBody{
			// 	TokenAddress:     utilityTokenAddress,
			// 	MinterPrivateKey: minterPrivatekey,
			// 	MintAmount:       totalUtilityTokenAmount,
			// 	ToAddress:        metadata.AccountAddress,
			// 	ChainKey:         metadata.ChainKey,
			// 	Network:          metadata.Network,
			// })
			// if err != nil {
			// 	logger.Error(err.Error())
			// 	return err
			// }
			return nil
		}()
	}
	wg.Wait()

	return nil
}
