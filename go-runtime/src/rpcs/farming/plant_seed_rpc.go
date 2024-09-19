package rpcs_farming

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_crops "cifarm-server/src/collections/crops"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type PlantSeedRpcParams struct {
	InventorySeedKey  string `json:"inventorySeedKey"`
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

type PlantSeedRpcResponse struct {
	HarvestIn int64 `json:"harvestIn"`
}

func PlantSeedRpc(
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

	var params *PlantSeedRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_inventories.ReadByKey(ctx, logger, db, nk, collections_inventories.ReadByKeyParams{
		Key:    params.InventorySeedKey,
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
	if inventory.Type != collections_inventories.TYPE_SEED {
		errMsg := "inventory not plant seed"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	object, err = collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemTileKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "object cannot be nil"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	placedItem, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if placedItem.Type != collections_placed_items.TYPE_TILE {
		errMsg := "placed item not farming tile"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if placedItem.IsPlanted {
		errMsg := "tile is planted"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	err = collections_inventories.Delete(ctx, logger, db, nk, collections_inventories.DeleteParams{
		Key:      params.InventorySeedKey,
		Quantity: 1,
		UserId:   userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err = collections_crops.ReadByKey(ctx, logger, db, nk, collections_crops.ReadByKeyParams{
		Key: inventory.ReferenceKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	crop, err := collections_common.ToValue[collections_crops.Crop](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if crop == nil {
		errMsg := "crop not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	placedItem.SeedGrowthInfo = collections_placed_items.SeedGrowthInfo{
		CurrentStage:             1,
		CurrentStageTimeElapsed:  0,
		TotalTimeElapsed:         0,
		HarvestQuantityRemaining: crop.MaxHarvestQuantity,
		Crop:                     *crop,
	}
	placedItem.IsPlanted = true
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *placedItem,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		UserId: userId,
		Amount: collections_config.EXPERIENCE_FROM_ACTIVITY,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(PlantSeedRpcResponse{HarvestIn: crop.GrowthStageDuration})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
