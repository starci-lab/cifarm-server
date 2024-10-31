package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HelpCureAnimalRpcParams struct {
	PlacedItemAnimalKey string `json:"placedItemAnimalKey"`
	UserId              string `json:"userId"`
}

func HelpCureAnimalRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *HelpCureAnimalRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//get activities
	object, err := collections_system.ReadActivities(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	activities, err := collections_common.ToValue[collections_system.Activities](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err = collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemAnimalKey,
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "animal not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	animal, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if animal.AnimalInfo.CurrentState != collections_placed_items.ANIMAL_CURRENT_STATE_SICK {
		errMsg := "animal is not sick"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process - ok
	//pay energy first, if not revert
	err = collections_player.DecreaseEnergy(ctx, logger, db, nk, collections_player.DecreaseEnergyParams{
		UserId: userId,
		Amount: activities.UseHerbicide.ExperiencesGain,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update the animal
	animal.AnimalInfo.CurrentState = collections_placed_items.ANIMAL_CURRENT_STATE_NORMAL

	//update the animal
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *animal,
		UserId:     params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//increase experience
	err = collections_player.IncreaseExperiences(ctx, logger, db, nk, collections_player.IncreaseExperiencesParams{
		UserId: userId,
		Amount: activities.HelpCureAnimal.ExperiencesGain,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
