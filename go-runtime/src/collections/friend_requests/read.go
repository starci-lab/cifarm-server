package collections_friend_requests

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
	Type         int    `json:"type"`
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
	query := fmt.Sprintf("+user_id:%s +value.referenceKey:%s +value.type:%v +value.premium:%s",
		params.UserId, params.ReferenceKey, params.Type, utils.BoolToStorageQuery(params.Premium))
	logger.Info(query)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	logger.Debug("%v", len(objects.Objects))
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

type ReadByTokenIdParams struct {
	TokenId      string `json:"tokenId"`
	ReferenceKey string `json:"referenceKey"`
}

func ReadByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByTokenIdParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_BY_TOKEN_ID
	logger.Info(fmt.Sprintf("+value.tokenId:%s +value.referenceKey:%s", params.TokenId, params.ReferenceKey))
	query := fmt.Sprintf("+value.tokenId:%s +value.referenceKey:%s", params.TokenId, params.ReferenceKey)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	logger.Debug("%v", len(objects.Objects))
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

type ReadManyUniqueParams struct {
	UserId       string `json:"userId"`
	ReferenceKey string `json:"referenceKey"`
}

func ReadManyUnique(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyUniqueParams,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_UNIQUE
	query := fmt.Sprintf("+user_id:%s +value.referenceKey:%s +value.unique:T", params.UserId, params.ReferenceKey)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, collections_common.MAX_ENTRIES, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

type ReadManyAvailableParams struct {
	UserId string `json:"userId"`
}

func ReadManyAvailable(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyAvailableParams,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_AVAILABLE
	query := fmt.Sprintf("+user_id:%s +value.isPlaced:F", params.UserId)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, collections_common.MAX_ENTRIES, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

type ReadManyParams struct {
	UserId string   `json:"userId"`
	Keys   []string `json:"keys"`
}

func ReadMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyParams,
) ([]*api.StorageObject, error) {
	var reads []*runtime.StorageRead
	for _, key := range params.Keys {
		read := runtime.StorageRead{
			Collection: COLLECTION_NAME,
			Key:        key,
			UserID:     params.UserId,
		}
		reads = append(reads, &read)
	}

	objects, err := nk.StorageRead(ctx, reads)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

type ReadManyDeliveringParams struct {
	UserId string `json:"userId"`
}
