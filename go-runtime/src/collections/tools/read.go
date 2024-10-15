package collections_tools

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadByKeyParams struct {
	Key string `json:"key"`
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
	objects, _, err := nk.StorageList(ctx, "", params.UserId, COLLECTION_NAME, collections_common.MAX_ENTRIES, "")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}
