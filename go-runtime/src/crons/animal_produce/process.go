package crons_animal_produce

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"math/rand"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ExecuteProcedureLogicParams struct {
	PlacedItem    *collections_placed_items.PlacedItem
	TimeInSeconds int64
	UserId        string
}

func ExecuteProcedureLogic(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params ExecuteProcedureLogicParams,
) error {
	//hungy, do nothing
	if params.PlacedItem.AnimalInfo.CurrentState == collections_placed_items.ANIMAL_CURRENT_STATE_HUNGRY {
		//need fed, do nothing
		return nil
	}
	if params.PlacedItem.AnimalInfo.HasYielded {
		//has yielded, do nothing
		return nil
	}

	if !params.PlacedItem.AnimalInfo.IsAdult {
		//do non-aldult logic
		//since this state need feed to grow
		//increase time hungry
		params.PlacedItem.AnimalInfo.CurrentHungryTime += params.TimeInSeconds
		params.PlacedItem.AnimalInfo.CurrentGrowthTime += params.TimeInSeconds

		turnIntoAldult := false
		if params.PlacedItem.AnimalInfo.CurrentGrowthTime >= params.PlacedItem.AnimalInfo.Animal.GrowthTime {
			//turn into adult logic
			params.PlacedItem.AnimalInfo.IsAdult = true
			turnIntoAldult = true
		}

		//if turn into adult or hungry, set state to hungry
		if turnIntoAldult || params.PlacedItem.AnimalInfo.CurrentHungryTime >= params.PlacedItem.AnimalInfo.Animal.HungerTime {
			params.PlacedItem.AnimalInfo.CurrentState = collections_placed_items.ANIMAL_CURRENT_STATE_HUNGRY
		}
	} else {
		//do adult logic
		//since if hungy returned, so that we do no hungry logic here
		params.PlacedItem.AnimalInfo.CurrentYieldTime += params.TimeInSeconds
		//sick here, sick happen randomly when animal is adult, and only one time
		if !params.PlacedItem.AnimalInfo.AlreadySick {
			randomValue := rand.Float64()
			if randomValue <= params.PlacedItem.AnimalInfo.Animal.SickChance {
				params.PlacedItem.AnimalInfo.CurrentState = collections_placed_items.ANIMAL_CURRENT_STATE_SICK
				params.PlacedItem.AnimalInfo.AlreadySick = true
			}
		}

		if params.PlacedItem.AnimalInfo.CurrentYieldTime >= params.PlacedItem.AnimalInfo.Animal.YieldTime {
			params.PlacedItem.AnimalInfo.CurrentYieldTime = 0
			params.PlacedItem.AnimalInfo.HasYielded = true
			if params.PlacedItem.AnimalInfo.CurrentState == collections_placed_items.ANIMAL_CURRENT_STATE_SICK {
				params.PlacedItem.AnimalInfo.HarvestQuantityRemaining = params.PlacedItem.AnimalInfo.Animal.MaxHarvestQuantity
			} else {
				params.PlacedItem.AnimalInfo.HarvestQuantityRemaining = (params.PlacedItem.AnimalInfo.Animal.MaxHarvestQuantity + params.PlacedItem.AnimalInfo.Animal.MinHarvestQuantity) / 2
			}
		}
	}

	_, err := collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *params.PlacedItem,
		UserId:     params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type HandleAnimalProcedureParams struct {
	UserId        string `json:"userId"`
	TimeInSeconds int64  `json:"timeInSeconds"`
}

func HandleAnimalProcedure(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params HandleAnimalProcedureParams,

) error {
	objects, err := collections_placed_items.ReadByFilters2(ctx, logger, db, nk, collections_placed_items.ReadByFilters2Params{
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
			err = ExecuteProcedureLogic(ctx, logger, db, nk, ExecuteProcedureLogicParams{
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
		speedUpTime = speedUp.AnimalProcedureTime
	}
	if speedUpTime > 0 {
		err := collections_system.WriteSpeedUp(ctx, logger, db, nk, collections_system.WriteSpeedUpParams{
			SpeedUp: collections_system.SpeedUp{
				AnimalProcedureTime: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	for _, userId := range users.UserIds {
		go HandleAnimalProcedure(ctx, logger, db, nk, HandleAnimalProcedureParams{
			UserId:        userId,
			TimeInSeconds: 1 + timeSinceLastUptime + speedUpTime,
		})
	}

	return err
}
