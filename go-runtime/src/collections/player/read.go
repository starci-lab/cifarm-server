package collections_player

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadMetadataParams struct {
	UserId string `json:"userId"`
}

func ReadMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadMetadataParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        KEY_METADATA,
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

type GetMetadataParams struct {
	Metadata Metadata `json:"metadata"`
}

func GetMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetMetadataParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_METADATA
	query := fmt.Sprintf("+value.accountAddress:%s +value.chainKey:%v +value.network:%s",
		params.Metadata.AccountAddress,
		params.Metadata.ChainKey,
		params.Metadata.Network,
	)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		errMsg := "metadata not found"
		logger.Error(errMsg)
		return nil, nil
	}
	return objects.Objects[0], nil
}

type GetUserIdByMetadataParams struct {
	Metadata Metadata `json:"metadata"`
}

func GetUserIdByMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetUserIdByMetadataParams,
) (string, error) {
	name := STORAGE_INDEX_USER_ID
	query := fmt.Sprintf(
		"+value.accountAddress:%s +value.network:%s +value.chainKey:%s",
		params.Metadata.AccountAddress,
		params.Metadata.Network,
		params.Metadata.ChainKey)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if len(objects.Objects) == 0 {
		return "", nil
	}
	var result = objects.Objects[0]
	return result.UserId, nil
}

type ReadMetadatasParams struct {
	TelegramUserId string `json:"telegramUserId"`
}

func ReadMetadatas(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadMetadatasParams,
) ([]*api.StorageObject, error) {
	name := STORAGE_INDEX_METADATAS
	query := fmt.Sprintf(
		"+value.telegramData.userId:%s",
		params.TelegramUserId)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, collections_common.MAX_ENTRIES_LIST, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects.Objects, nil
}

type ReadVisitStateParams struct {
	UserId string `json:"userId"`
}

func ReadVisitState(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadVisitStateParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        KEY_VISIT_STATE,
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

type ReadPlayerStatsParams struct {
	UserId string `json:"userId"`
}

func ReadPlayerStats(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadPlayerStatsParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        KEY_PLAYER_STATS,
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

type ReadRewardTrackerParams struct {
	UserId string `json:"userId"`
}

func ReadRewardTracker(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadRewardTrackerParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        KEY_REWARD_TRACKER,
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

type ReadFollowingsParams struct {
	UserId string `json:"userId"`
}

func ReadFollowings(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadFollowingsParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        KEY_FOLLOWINGS,
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
