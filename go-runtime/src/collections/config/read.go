package collections_config

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadMetadataByKeyParams struct {
	UserId string `json:"userId"`
}

func ReadMetadataByKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadMetadataByKeyParams,
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
		`+value.accountAddress:%s +value.network:%s +value.chainKey:%s`,
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
