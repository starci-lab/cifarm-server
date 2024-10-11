package crons_seed_growth

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"math/rand/v2"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ExecuteGrowthLogicParams struct {
	PlacedItem    *collections_placed_items.PlacedItem
	TimeInSeconds int64
	UserId        string
}

func ExecuteGrowthLogic(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params ExecuteGrowthLogicParams,
) error {
	if !params.PlacedItem.SeedGrowthInfo.IsPlanted {
		//do nothing via nothing planted
		return nil
	}
	if params.PlacedItem.SeedGrowthInfo.FullyMatured {
		//do nothing via fully matured
		return nil
	}

	if params.PlacedItem.SeedGrowthInfo.CurrentState == collections_placed_items.CURRENT_STATE_NEED_WATER {
		//do nothing via being need watering
		return nil
	}

	object, err := collections_system.ReadGlobalConstants(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "global constants not found"
		logger.Error(errMsg)
		return err
	}
	globalConstants, err := collections_common.ToValue[collections_system.GlobalConstants](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	params.PlacedItem.SeedGrowthInfo.TotalTimeElapsed += params.TimeInSeconds
	params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed += params.TimeInSeconds

	var loopCounter int
	for {
		loopCounter += 1
		if loopCounter > 5 {
			break
		}
		if params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed >= params.PlacedItem.SeedGrowthInfo.Crop.GrowthStageDuration {
			params.PlacedItem.SeedGrowthInfo.CurrentStageTimeElapsed -= params.PlacedItem.SeedGrowthInfo.Crop.GrowthStageDuration
			params.PlacedItem.SeedGrowthInfo.CurrentStage += 1
			//reset fertilized status after growth stage
			params.PlacedItem.SeedGrowthInfo.IsFertilized = false

			if params.PlacedItem.SeedGrowthInfo.CurrentStage <= 3 {
				//50% chance to be drain,
				if rand.Float64() < globalConstants.GameRandomness.NeedWater {
					params.PlacedItem.SeedGrowthInfo.CurrentState = collections_placed_items.CURRENT_STATE_NEED_WATER
				}
			}

			if params.PlacedItem.SeedGrowthInfo.CurrentStage == 4 {
				//50% to be infested or weedly, chance maybe difference via better tiles
				if rand.Float64() <= globalConstants.GameRandomness.IsWeedyOrInfested {
					if rand.Float64() < 0.5 {
						params.PlacedItem.SeedGrowthInfo.CurrentState = collections_placed_items.CURRENT_STATE_IS_WEEDY
					} else {
						params.PlacedItem.SeedGrowthInfo.CurrentState = collections_placed_items.CURRENT_STATE_IS_INFESTED
					}
				}
			}

			if params.PlacedItem.SeedGrowthInfo.CurrentStage == params.PlacedItem.SeedGrowthInfo.Crop.GrowthStages {
				if params.PlacedItem.SeedGrowthInfo.CurrentState == collections_placed_items.CURRENT_STATE_IS_WEEDY ||
					params.PlacedItem.SeedGrowthInfo.CurrentState == collections_placed_items.CURRENT_STATE_IS_INFESTED {
					//reduce quantity
					newQuantity := (params.PlacedItem.SeedGrowthInfo.Crop.MaxHarvestQuantity + params.PlacedItem.SeedGrowthInfo.Crop.MinHarvestQuantity) / 2
					params.PlacedItem.SeedGrowthInfo.HarvestQuantityRemaining = newQuantity
				}
				params.PlacedItem.SeedGrowthInfo.FullyMatured = true
				break
			}
		} else {
			break
		}
	}

	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *params.PlacedItem,
		UserId:     params.UserId,
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
	objects, err := collections_placed_items.ReadByFilters1(ctx, logger, db, nk, collections_placed_items.ReadByFilters1Params{
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
	var speedUpTime int64
	object, err = collections_system.ReadSpeedUp(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object != nil {
		speedUp, err := collections_common.ToValue[collections_system.SpeedUp](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		speedUpTime = speedUp.SeedGrowthTime
	}
	if speedUpTime > 0 {
		err := collections_system.WriteSpeedUp(ctx, logger, db, nk, collections_system.WriteSpeedUpParams{
			SpeedUp: collections_system.SpeedUp{
				SeedGrowthTime: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	for _, userId := range users.UserIds {
		go HandleSeedGrowth(ctx, logger, db, nk, HandleSeedGrowthParams{
			UserId:        userId,
			TimeInSeconds: 1 + timeSinceLastUptime + speedUpTime,
		})
	}

	return err
}
