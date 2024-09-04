package placed_items

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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
