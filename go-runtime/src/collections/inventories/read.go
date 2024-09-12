package collections_inventories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadByReferenceKeyParams struct {
	ReferenceKey string `json:"referenceKey"`
	UserId       string `json:"userId"`
}

func ReadByReferenceKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByReferenceKeyParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_BY_REFERENCE_KEY
	query := fmt.Sprintf(`+value.referenceKey:%s`, params.ReferenceKey)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, params.UserId, name, query, 1, order)
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
	var object = objects[0]
	return object, nil
}
