package collections_placed_items

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadsParams struct {
	UserId string `json:"userId"`
}

func ReadMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadsParams,
) ([]*api.StorageObject, error) {
	objects, _, err := nk.StorageList(ctx, params.UserId, params.UserId, COLLECTION_NAME, collections_common.MAX_ENTRIES, "")
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
	params ReadByKeyParams,
) (*api.StorageObjects, error) {

	name := STORAGE_INDEX_BY_FILTERS_1
	query := fmt.Sprintf(`+value.isPlanted:T -fullyMatured:T +value.type:%v`, TYPE_TILE)
	maxEntries := collections_common.MAX_ENTRIES
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, params.UserId, name, query, maxEntries, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, err
}
