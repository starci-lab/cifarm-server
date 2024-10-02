package rpcs_farming

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type CollectAnimalProductRpcParams struct {
	PlacedItemAnimalKey string `json:"placedItemAnimalKey"`
}

type CollectAnimalProductRpcResponse struct {
	InventoryAnimalProductKey string `json:"inventoryAnimalProductKey"`
}

func CollectAnimalProductRpc(
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

	var params *CollectAnimalProductRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

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
	//not yielded
	if !animal.AnimalInfo.HasYielded {
		errMsg := "animal has not yielded"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//update status
	animal.AnimalInfo.HasYielded = false
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *animal,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//create inventory
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: animal.Key,
			Quantity:     animal.AnimalInfo.HarvestQuantityRemaining,
			Type:         collections_inventories.TYPE_ANIMAL_PRODUCT,
			Deliverable:  true,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//exp
	var experiences int64
	if animal.AnimalInfo.Animal.IsNFT {
		experiences = animal.SeedGrowthInfo.Crop.PremiumHarvestExperiences
	} else {
		experiences = animal.SeedGrowthInfo.Crop.BasicHarvestExperiences
	}

	err = collections_config.IncreaseExperiences(ctx, logger, db, nk, collections_config.IncreaseExperiencesParams{
		Amount: experiences,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(CollectAnimalProductRpcResponse{
		InventoryAnimalProductKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
