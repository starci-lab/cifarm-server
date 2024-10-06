package collections_delivering_products

import (
	collections_common "cifarm-server/src/collections/common"
	"cifarm-server/src/utils"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadParams struct {
	ReferenceKey string `json:"referenceKey"`
	UserId       string `json:"userId"`
	Index        int    `json:"index"`
	Premium      bool   `json:"premium"`
}

func Read(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX
	query := fmt.Sprintf("+user_id:%s +value.referenceKey:%s +value.index:%v +value.isPremium:%s",
		params.UserId,
		params.ReferenceKey,
		params.Index,
		utils.BoolToStorageQuery(params.Premium))
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
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

type ReadByTokenIdParams struct {
	TokenId      int    `json:"tokenId"`
	ReferenceKey string `json:"referenceKey"`
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
	objects, _, err := nk.StorageList(ctx, params.UserId, params.UserId, COLLECTION_NAME, collections_common.MAX_ENTRIES_LIST, "")
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}
