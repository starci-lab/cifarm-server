package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_supplies "cifarm-server/src/collections/supplies"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HelpUseFertilizerRpcParams struct {
	InventoryFertilizerKey string `json:"inventoryFertilizerKey"`
	PlacedItemTileKey      string `json:"placedItemTileKey"`
	UserId                 string `json:"userId"`
}

func HelpUseFertilizerRpc(
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

	var params *HelpUseFertilizerRpcParams
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

	object, err = collections_inventories.ReadByKey(ctx, logger, db, nk, collections_inventories.ReadByKeyParams{
		Key:    params.InventoryFertilizerKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	inventory, err := collections_common.ToValue[collections_inventories.Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if inventory.Type != collections_inventories.TYPE_SUPPLY {
		errMsg := "invalid inventory type"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//reference to the supply
	object, err = collections_supplies.ReadByKey(ctx, logger, db, nk, collections_supplies.ReadByKeyParams{
		Key: inventory.ReferenceKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "supply not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	supply, err := collections_common.ToValue[collections_supplies.Supply](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if supply.Type != collections_supplies.TYPE_FERTILIZER {
		errMsg := "invalid supply type"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//reducer timer speed
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
	if tile.SeedGrowthInfo.IsFertilized {
		errMsg := "tile already fertilized"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process - ok
	//pay energy first, if not revert
	err = collections_config.DecreaseEnergy(ctx, logger, db, nk, collections_config.DecreaseEnergyParams{
		UserId: userId,
		Amount: activities.HelpUseFertilizer.ExperiencesGain,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//delete inventory
	err = collections_inventories.Delete(ctx, logger, db, nk, collections_inventories.DeleteParams{
		Key:      params.InventoryFertilizerKey,
		Quantity: 1,
		UserId:   userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update tile
	tile.SeedGrowthInfo.IsFertilized = true
	tile.SeedGrowthInfo.CurrentStageTimeElapsed += supply.FertilizerEffect.TimeReduce
	tile.SeedGrowthInfo.TotalTimeElapsed += supply.FertilizerEffect.TimeReduce

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

	//gain experiences
	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		UserId: userId,
		Amount: activities.HelpUseFertilizer.ExperiencesGain * multiplier,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
