package crons_seed_growth

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HandleSeedGrowthParams struct {
	UserId string `json:"userId"`
}

func HandleSeedGrowth(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params HandleSeedGrowthParams,
) error {
	objects, err := collections_placed_items.ReadByFilters1(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.UserId,
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for _, object := range objects.Objects {
		go func() error {
			placedItem, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			if placedItem.SeedGrowthInfo.CurrentStage == placedItem.SeedGrowthInfo.Seed.GrowthStages {
				return nil
			}
			time := int64(1)
			placedItem.SeedGrowthInfo.TotalTimeElapsed += time
			placedItem.SeedGrowthInfo.CurrentStageTimeElapsed += time

			if placedItem.SeedGrowthInfo.CurrentStageTimeElapsed >= placedItem.SeedGrowthInfo.Seed.GrowthStageDuration {
				placedItem.SeedGrowthInfo.CurrentStageTimeElapsed -= placedItem.SeedGrowthInfo.Seed.GrowthStageDuration
				placedItem.SeedGrowthInfo.CurrentStage += 1
			}
			err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
				PlacedItem: *placedItem,
				UserId:     params.UserId,
				Key:        object.Key,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			return nil
		}()
	}
	return nil
}

func Process(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	object, err := collections_system.ReadByKey(ctx, logger, db, nk)
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
		go HandleSeedGrowth(ctx, logger, db, nk, HandleSeedGrowthParams{
			UserId: userId,
		})
	}

	return err
}
