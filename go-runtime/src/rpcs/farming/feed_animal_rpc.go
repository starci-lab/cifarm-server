package rpcs_farming

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type FeedAnimalRpcParams struct {
	PlacedItemAnimalKey    string `json:"placedItemAnimalKey"`
	InventoryAnimalFeedKey string `json:"inventoryAnimalFeedKey"`
}

func FeedAnimalRpc(
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

	var params *FeedAnimalRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//fetch the animal
	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemAnimalKey,
		UserId: userId,
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

	if animal.Type != collections_placed_items.TYPE_ANIMAL {
		errMsg := "placed item is not an animal"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !animal.AnimalInfo.NeedFed {
		errMsg := "animal does not need to be fed"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//delete the feed from inventory
	err = collections_inventories.Delete(ctx, logger, db, nk, collections_inventories.DeleteParams{
		Key:      params.InventoryAnimalFeedKey,
		Quantity: 1,
		UserId:   userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update animal status
	animal.AnimalInfo.NeedFed = false

	//update the animal
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *animal,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//increase user experience
	object, err = collections_system.ReadActivityExperiences(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	activityExperiences, err := collections_common.ToValue[collections_system.ActivityExperiences](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		UserId: userId,
		Amount: activityExperiences.UseFertilizer,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
