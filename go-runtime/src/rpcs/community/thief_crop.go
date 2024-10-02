package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/utils"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math"
	"math/rand"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ThiefCropRpcParams struct {
	UserId            string `json:"userId"`
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

type ThiefCropRpcResponse struct {
	ThiefCropInventoryKey string `json:"thiefCropInventoryKey"`
	ThiefQuantity         int    `json:"thiefQuantity"`
}

func ThiefCropRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *ThiefCropRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if userId == params.UserId {
		errMsg := "you cannot theif your plants"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	// check, err := friends.CheckFriendByUserId(ctx, logger, db, nk, friends.CheckFriendByUserIdParams{
	// 	UserId:       userId,
	// 	FriendUserId: params.UserId,
	// })
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return "", err
	// }

	// if !check {
	// 	errMsg := "not your friend"
	// 	logger.Error(errMsg)
	// 	return "", errors.New(errMsg)
	// }

	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
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

	if !tile.SeedGrowthInfo.FullyMatured {
		errMsg := "crop not fully matured"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if utils.Contains(tile.SeedGrowthInfo.ThiefedBy, userId) {
		errMsg := "theif the crop before"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	maximunTheifQuantity := tile.SeedGrowthInfo.HarvestQuantityRemaining - tile.SeedGrowthInfo.Crop.MinHarvestQuantity
	if maximunTheifQuantity == 0 {
		errMsg := "cannot thief anymore"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//fn to calculate
	thiefQuantity := 1
	random := rand.Float64()
	if random > 0.95 {
		thiefQuantity = 3
	} else if random > 0.8 {
		thiefQuantity = 2
	}
	thiefQuantity = int(math.Min(float64(maximunTheifQuantity), float64(thiefQuantity)))

	//check kinh nghiệm, check các thứ, ...
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: tile.SeedGrowthInfo.Crop.Key,
			Type:         collections_inventories.TYPE_HARVESTED_CROP,
			Quantity:     thiefQuantity,
			IsPremium:    tile.ReferenceKey == collections_tiles.KEY_PREMIUM,
			Deliverable:  true,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//giam san luong
	//add thang an trom vao list
	tile.SeedGrowthInfo.HarvestQuantityRemaining -= thiefQuantity
	tile.SeedGrowthInfo.ThiefedBy = append(tile.SeedGrowthInfo.ThiefedBy, userId)

	//update the tile
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *tile,
		UserId:     params.UserId,
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

	value, err := json.Marshal(ThiefCropRpcResponse{
		ThiefCropInventoryKey: result.Key,
		ThiefQuantity:         thiefQuantity,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
