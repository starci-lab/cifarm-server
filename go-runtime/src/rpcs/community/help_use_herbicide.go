package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HelpUseHerbicideRpcParams struct {
	UserId            string `json:"userId"`
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

func HelpUseHerbicideRpc(
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

	var params *HelpUseHerbicideRpcParams
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

	if userId == params.UserId {
		errMsg := "you cannot help yourself with using herbicide"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	object, err = collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemTileKey,
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "tile not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	tile, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if !tile.SeedGrowthInfo.IsPlanted {
		errMsg := "tile is not being planted"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if tile.SeedGrowthInfo.PlantCurrentState != collections_placed_items.PLANT_CURRENT_STATE_IS_WEEDY {
		errMsg := "plant is not weedy"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process - ok
	//pay energy first, if not revert
	err = collections_config.DecreaseEnergy(ctx, logger, db, nk, collections_config.DecreaseEnergyParams{
		UserId: userId,
		Amount: activities.HelpUseHerbicide.EnergyCost,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update tile status
	tile.SeedGrowthInfo.PlantCurrentState = collections_placed_items.PLANT_CURRENT_STATE_NORMAL

	//update the tile
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *tile,
		UserId:     params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check friend
	multiplier, err := GetMutipleValue(ctx, logger, db, nk, GetMutipleValueParams{
		UserId:      userId,
		OtherUserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//increase experience
	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		UserId: userId,
		Amount: activities.HelpUseHerbicide.ExperiencesGain * multiplier,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
