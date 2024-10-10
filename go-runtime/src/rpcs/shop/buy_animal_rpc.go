package rpcs_shop

import (
	collections_animals "cifarm-server/src/collections/animals"
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuyAnimalRpcParams struct {
	Key                   string `json:"key"`
	PlacedItemBuildingKey string `json:"placedItemBuildingKey"`
}

type BuyAnimalRpcResponse struct {
	PlacedItemAnimalKey string `json:"placedItemAnimalKey"`
}

func BuyAnimalRpc(ctx context.Context,
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

	var params *BuyAnimalRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_animals.ReadByKey(ctx, logger, db, nk, collections_animals.ReadByKeyParams{
		Key: params.Key,
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

	animal, err := collections_common.ToValue[collections_animals.Animal](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if animal == nil {
		errMsg := "animal not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if !animal.AvailableInShop {
		errMsg := "not available in shop"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//check the parent building
	object, err = collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemBuildingKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "parent building not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	building, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check whether the parent building is a building
	if building.Type != collections_placed_items.TYPE_BUILDING {
		errMsg := "parent building is not a building"
		logger.Error(errMsg)
		return "", err
	}
	//check animal type
	if building.BuildingInfo.Building.AnimalKey != animal.Key {
		errMsg := "animal type does not match"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	//check ocupation
	if building.BuildingInfo.Occupancy >= building.BuildingInfo.Building.Capacity {
		errMsg := "building is full"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//reduce money
	err = _wallets.UpdateWalletGolds(ctx, logger, db, nk, _wallets.UpdateWalletGoldsParams{
		Amount: -animal.OffspringPrice,
		UserId: userId,
		Metadata: map[string]interface{}{
			"name": "Buy animal",
			"key":  params.Key,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	// we place animal in the parent building
	result, err := collections_placed_items.Write(ctx,
		logger, db, nk,
		collections_placed_items.WriteParams{
			PlacedItem: collections_placed_items.PlacedItem{
				ReferenceKey:        params.Key,
				ParentPlacedItemKey: building.Key,
				AnimalInfo: collections_placed_items.AnimalInfo{
					Animal: *animal,
				},
				Type: collections_inventories.TYPE_ANIMAL,
			},
			UserId: userId,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	//update occupancy
	building.BuildingInfo.Occupancy++
	_, err = collections_placed_items.Write(ctx,
		logger, db, nk,
		collections_placed_items.WriteParams{
			PlacedItem: *building,
			UserId:     userId,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuyAnimalRpcResponse{
		PlacedItemAnimalKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
