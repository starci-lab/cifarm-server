package collections_inventories

import (
	collections_common "cifarm-server/src/collections/common"
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
	query := fmt.Sprintf("+user_id:%s +value.referenceKey:%s", params.UserId, params.ReferenceKey)
	logger.Warn("+user_id:%s +value.referenceKey:%s", params.UserId, params.ReferenceKey)
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

func ReadByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByTokenIdParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_BY_TOKEN_ID
	query := fmt.Sprintf("+value.tokenId:%v +value.referenceKey:%s", params.TokenId, params.ReferenceKey)
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

type ReadManyByUserIdParams struct {
	UserId       string `json:"userId"`
	ReferenceKey string `json:"referenceKey"`
}

func ReadManyByUserId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyByUserIdParams,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_BY_USER_ID
	logger.Info("+user_id:%s +value.referenceKey:%s", params.UserId, params.ReferenceKey)
	query := fmt.Sprintf("+user_id:%s +value.referenceKey:%s", params.UserId, params.ReferenceKey)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, collections_common.MAX_ENTRIES, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}

type ReadManyByUserIdNonPlacedParams struct {
	UserId string `json:"userId"`
}

func ReadManyByUserIdNonPlaced(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadManyByUserIdNonPlacedParams,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_BY_USER_ID_NON_PLACED
	query := fmt.Sprintf("+user_id:%s +value.isPlaced:F", params.UserId)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, collections_common.MAX_ENTRIES, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}
