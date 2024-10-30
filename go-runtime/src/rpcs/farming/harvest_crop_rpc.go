package rpcs_farming

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	collections_tiles "cifarm-server/src/collections/tiles"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HarvestCropRpcParams struct {
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

type HarvestCropRpcResponse struct {
	InventoryHarvestedCropKey string `json:"inventoryHarvestedCropKey"`
}

func HarvestCropRpc(
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

	var params *HarvestCropRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemTileKey,
		UserId: userId,
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

	value1, err := json.Marshal(tile)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	logger.Info(string(value1))

	if !tile.SeedGrowthInfo.IsPlanted {
		errMsg := "tile is not being planted"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !tile.SeedGrowthInfo.FullyMatured {
		errMsg := "plant not fully matured"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//check if the crop is premium
	premium := tile.ReferenceKey == collections_tiles.KEY_FERTILE_TILE
	//write to inventories the havested items
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: tile.SeedGrowthInfo.Crop.Key,
			Type:         collections_inventories.TYPE_HARVESTED_CROP,
			Quantity:     tile.SeedGrowthInfo.HarvestQuantityRemaining,
			Premium:      premium,
			Deliverable:  true,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var experiences int
	if premium {
		experiences = tile.SeedGrowthInfo.Crop.PremiumHarvestExperiences
	} else {
		experiences = tile.SeedGrowthInfo.Crop.BasicHarvestExperiences
	}

	err = collections_player.IncreaseExperiences(ctx, logger, db, nk, collections_player.IncreaseExperiencesParams{
		Amount: experiences,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update tile status
	tile.SeedGrowthInfo.FullyMatured = false
	tile.SeedGrowthInfo.IsPlanted = false
	tile.SeedGrowthInfo = collections_placed_items.SeedGrowthInfo{}

	//update the tile
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *tile,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(HarvestCropRpcResponse{
		InventoryHarvestedCropKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
