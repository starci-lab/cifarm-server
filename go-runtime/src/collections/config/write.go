package collections_config

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteMetadataParams struct {
	Metadata Metadata `json:"metadata"`
	UserId   string   `json:"userId"`
}

func WriteMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteMetadataParams,
) error {
	value, err := json.Marshal(params.Metadata)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_METADATA,
			UserID:          params.UserId,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type WriteVisitStateParams struct {
	VisitState VisitState `json:"visitState"`
	UserId     string     `json:"userId"`
}

func WriteVisitState(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteVisitStateParams,
) error {
	value, err := json.Marshal(params.VisitState)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_VISIT_STATE,
			UserID:          params.UserId,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
