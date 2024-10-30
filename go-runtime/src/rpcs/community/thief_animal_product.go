package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/utils"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ThiefAnimalProductRpcParams struct {
	UserId              string `json:"userId"`
	PlacedItemAnimalKey string `json:"placedItemAnimalKey"`
}

type ThiefAnimalProductRpcResponse struct {
	InventoryThiefCropKey string `json:"inventoryThiefCropKey"`
	ThiefQuantity         int    `json:"thiefQuantity"`
}

func ThiefAnimalProductRpc(
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

	var params *ThiefAnimalProductRpcParams
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
		errMsg := "you cannot theif your animals"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//ensure you have more level
	//your level
	object, err = collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	playerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//other level
	object, err = collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "the other's player stats not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	otherPlayerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check level
	if playerStats.LevelInfo.Level < otherPlayerStats.LevelInfo.Level {
		errMsg := "you cannot theif higher level"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
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
		errMsg := "tile not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	animal, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if !animal.AnimalInfo.IsAdult {
		errMsg := "animal is not adult"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !animal.AnimalInfo.HasYielded {
		errMsg := "animal has not yielded"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if utils.Contains(animal.AnimalInfo.ThiefedBy, userId) {
		errMsg := "theif the animal before"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	maximunTheifQuantity := animal.AnimalInfo.HarvestQuantityRemaining - animal.AnimalInfo.Animal.MinHarvestQuantity
	if maximunTheifQuantity == 0 {
		errMsg := "cannot thief anymore"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process - ok
	//pay energy first, if not revert
	err = collections_player.DecreaseEnergy(ctx, logger, db, nk, collections_player.DecreaseEnergyParams{
		UserId: userId,
		Amount: activities.ThiefAnimalProduct.EnergyCost,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//fn to calculate
	thiefQuantity, err := GetThiefValue(ctx, logger, db, nk, GetThiefValueParams{
		MaximunTheifQuantity: maximunTheifQuantity,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check kinh nghiệm, check các thứ, ...
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: animal.AnimalInfo.Animal.Key,
			Type:         collections_inventories.TYPE_HARVESTED_CROP,
			Quantity:     thiefQuantity,
			Premium:      animal.AnimalInfo.Animal.IsNFT,
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
	animal.AnimalInfo.HarvestQuantityRemaining -= thiefQuantity
	animal.AnimalInfo.ThiefedBy = append(animal.AnimalInfo.ThiefedBy, userId)

	//update the tile
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
		Amount: activities.ThiefAnimalProduct.ExperiencesGain,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(ThiefCropRpcResponse{
		InventoryThiefCropKey: result.Key,
		ThiefQuantity:         thiefQuantity,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
