package crons_seed_growth

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ExecuteGrowthLogicParams struct {
	PlacedItem    *collections_placed_items.PlacedItem
	TimeInSeconds int64
	UserId        string
	Key           string
}

func ExecuteGrowthLogic(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params ExecuteGrowthLogicParams,
) error {
	if params.PlacedItem.FullyMatured {
		return nil
	}
	params.PlacedItem.SeedGrowthInfo.TotalTimeElapsed += params.TimeInSeconds
	params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed += params.TimeInSeconds

	var loopCounter int
	for {
		loopCounter += 1
		if loopCounter > 5 {
			break
		}
		if params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed >= params.PlacedItem.SeedGrowthInfo.Seed.GrowthStageDuration {
			params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed -= params.PlacedItem.SeedGrowthInfo.Seed.GrowthStageDuration
			params.PlacedItem.SeedGrowthInfo.CurrentStage += 1
			if params.PlacedItem.SeedGrowthInfo.CurrentStage == params.PlacedItem.SeedGrowthInfo.Seed.GrowthStages {
				params.PlacedItem.FullyMatured = true
				break
			}
		} else {
			break
		}
	}

	err := collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *params.PlacedItem,
		UserId:     params.UserId,
		Key:        params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type HandleSeedGrowthParams struct {
	UserId        string `json:"userId"`
	TimeInSeconds int64  `json:"timeInSeconds"`
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
			err = ExecuteGrowthLogic(ctx, logger, db, nk, ExecuteGrowthLogicParams{
				PlacedItem:    placedItem,
				TimeInSeconds: params.TimeInSeconds,
				UserId:        params.UserId,
				Key:           object.Key,
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
	timeSinceLastUptime int64,
) error {
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
		go HandleSeedGrowth(ctx, logger, db, nk, HandleSeedGrowthParams{
			UserId:        userId,
			TimeInSeconds: 1 + timeSinceLastUptime,
		})
	}

	return err
}
