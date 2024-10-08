package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/friends"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HelpUsePestisideRpcParams struct {
	UserId            string `json:"userId"`
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

func HelpUsePestisideRpc(
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

	var params *HelpUsePestisideRpcParams
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

	if tile.SeedGrowthInfo.PlantCurrentState != collections_placed_items.PLANT_CURRENT_STATE_IS_INFESTED {
		errMsg := "plant is not infested"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process - ok
	//pay energy first, if not revert
	err = collections_config.DecreaseEnergy(ctx, logger, db, nk, collections_config.DecreaseEnergyParams{
		UserId: userId,
		Amount: activities.HelpUsePestiside.EnergyCost,
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
	check, err := friends.CheckFriendByUserId(ctx, logger, db, nk, friends.CheckFriendByUserIdParams{
		UserId:       userId,
		FriendUserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	multiplier := 1
	if check {
		multiplier = 2
	}

	//increase experience
	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		UserId: userId,
		Amount: activities.HelpUsePestiside.ExperiencesGain * multiplier,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
