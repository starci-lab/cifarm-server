package collections_nfts

import (
	"cifarm-server/src/utils"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

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
	TokenId  int    `json:"tokenId"`
	Type     int    `json:"type"`
	ChainKey string `json:"string"`
	Network  string `json:"network"`
}

func ReadByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByTokenIdParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_BY_TOKEN_ID
	query := fmt.Sprintf(`+value.tokenId:%v +value.type:%v +value.chainKey:%s +value.network:%s`,
		params.TokenId,
		params.Type,
		params.ChainKey,
		params.Network,
	)
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

type ReadByTokenIdsParams struct {
	TokenIds []int  `json:"tokenIds"`
	Type     int    `json:"type"`
	ChainKey string `json:"string"`
	Network  string `json:"network"`
}

func ReadByTokenIds(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByTokenIdsParams,
) (*api.StorageObjects, error) {
	name := STORAGE_INDEX_BY_TOKEN_ID
	query := fmt.Sprintf(`+value.tokenId:%v +value.type:%v +value.chainKey:%s +value.network:%s`,
		utils.SliceToString(params.TokenIds),
		params.Type,
		params.ChainKey,
		params.Network,
	)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return objects, nil
}
