package placed_items

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadPlacedItemObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*api.StorageObjects, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	name := _constants.STORAGE_INDEX_PLACED_ITEMS
	query := ""
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, userId, name, query, 10000, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

func ReadPlacedItemObjectsTypeFarmingTileIsPlanted(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	userId string,
) (*api.StorageObjects, error) {
	if userId == "" {
		_userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
		if !ok {
			errMsg := "user ID not found"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		userId = _userId
	}

	name := _constants.STORAGE_INDEX_PLACED_ITEMS
	query := fmt.Sprintf(`+value.type:%v +value.isPlanted:T`, _constants.PLACED_ITEM_TYPE_FARMING_TILE)
	maxEntries := math.MaxInt
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, userId, name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

func ReadPlacedItemObjectByKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	key string,
) (*api.StorageObject, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	name := _constants.STORAGE_INDEX_PLACED_ITEMS
	query := fmt.Sprintf(`+key:%s`, key)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}
	var object = objects.Objects[0]
	return object, nil
}

func ToPlacedItem(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.PlacedItem, error) {
	if object == nil {
		return nil, nil
	}
	var placedItem *_collections.PlacedItem
	err := json.Unmarshal([]byte(object.Value), &placedItem)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return placedItem, nil
}

func ToPlacedItems(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	objects *api.StorageObjects,
) (*[]_collections.PlacedItem, error) {

	placedItems := []_collections.PlacedItem{}

	for _, object := range objects.Objects {
		var placedItem *_collections.PlacedItem
		err := json.Unmarshal([]byte(object.Value), &placedItem)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		placedItems = append(placedItems, *placedItem)
	}
	return &placedItems, nil
}
