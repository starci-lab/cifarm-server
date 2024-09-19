package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/friends"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math"
	"math/rand"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ThiefPlantRpcParams struct {
	UserId            string `json:"userId"`
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

type ThiefPlantRpcResponse struct {
	TheifPlantInventoryKey string `json:"thiefPlantInventoryKey"`
}

func ThiefPlantRpc(
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

	var params *ThiefPlantRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	check, err := friends.CheckFriendByUserId(ctx, logger, db, nk, friends.CheckFriendByUserIdParams{
		UserId:       userId,
		FriendUserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if !check {
		errMsg := "not your friend"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
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

	if !tile.IsPlanted {
		errMsg := "tile is not being planted"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !tile.FullyMatured {
		errMsg := "plant not fully matured"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	maximunTheifQuantity := tile.SeedGrowthInfo.HarvestQuantityRemaining - tile.SeedGrowthInfo.Seed.MaxHarvestQuantity
	if maximunTheifQuantity == 0 {
		errMsg := "cannot thief anymore"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//fn to calculate
	theifQuantity := 1
	random := rand.Float64()
	if random > 0.95 {
		theifQuantity = 3
	} else if random > 0.8 {
		theifQuantity = 2
	}
	theifQuantity = int(math.Min(float64(maximunTheifQuantity), float64(theifQuantity)))

	//check kinh nghiệm, check các thứ, ...
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: tile.SeedGrowthInfo.Seed.Key,
			Type:         collections_inventories.TYPE_HARVESTED_PLANT,
			Quantity:     theifQuantity,
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
	tile.SeedGrowthInfo.HarvestQuantityRemaining -= theifQuantity
	//update the tile
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *tile,
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

	value, err := json.Marshal(ThiefPlantRpcResponse{
		TheifPlantInventoryKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
