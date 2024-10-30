package collections_placed_items

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadManyParams struct {
	UserId string `json:"userId"`
}

func ReadMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyParams,
) ([]*api.StorageObject, error) {
	objects, _, err := nk.StorageList(ctx, params.UserId, params.UserId, COLLECTION_NAME, collections_common.MAX_ENTRIES_LIST, "")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

type ReadByKeyParams struct {
	Key    string `json:"key"`
	UserId string `json:"userId"`
}

func ReadByKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByKeyParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        params.Key,
			UserID:     params.UserId,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects) == 0 {
		return nil, nil
	}

	object := objects[0]
	return object, nil
}

type ReadByFilters1Params struct {
	UserId string `json:"userId"`
}

func ReadByFilters1(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByFilters1Params,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_BY_FILTERS_1
	query := fmt.Sprintf("+user_id:%s +value.seedGrowthInfo.isPlanted:T -value.seedGrowthInfo.fullyMatured:T +value.type:%v", params.UserId, TYPE_TILE)
	maxEntries := collections_common.MAX_ENTRIES
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, err
}

type ReadByFilters2Params struct {
	UserId string `json:"userId"`
}

func ReadByFilters2(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByFilters2Params,
) (*api.StorageObjects, error) {

	name := STORAGE_INDEX_BY_FILTERS_2
	query := fmt.Sprintf("+user_id:%s +value.type:%v", params.UserId, TYPE_ANIMAL)
	maxEntries := collections_common.MAX_ENTRIES
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, err
}

type ReadByInventoryKeyParams struct {
	InventoryKey string `json:"inventoryKey"`
}

func ReadByInventoryKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByInventoryKeyParams,
) (*api.StorageObject, error) {

	name := STORAGE_INDEX_BY_INVENTORY_KEY
	query := fmt.Sprintf(`+value.inventoryKey:%s`, params.InventoryKey)
	maxEntries := 1
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}

	object := objects.Objects[0]
	return object, nil
}

type ReadByFilters3Params struct {
	UserId       string `json:"userId"`
	ReferenceKey string `json:"referenceKey"`
}

func ReadByFilters3(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByFilters3Params,
) (*api.StorageObjects, error) {

	name := STORAGE_INDEX_BY_FILTERS_3
	query := fmt.Sprintf("+user_id:%s +value.type:%v +value.referenceKey:%v", params.UserId, TYPE_TILE, params.ReferenceKey)

	maxEntries := collections_common.MAX_ENTRIES
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, err
}
