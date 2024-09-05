package collections_inventories

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	Inventory Inventory `json:"inventory"`
	UserId    string    `json:"userId"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) error {
	object, err := ReadByReferenceId(ctx, logger, db, nk, ReadByReferenceIdParams{
		ReferenceId: params.Inventory.ReferenceId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if object != nil {
		inventory, err := collections_common.ToValue[Inventory](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		inventory.Quantity += inventory.Quantity
		data, err := json.Marshal(inventory)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
			{
				Collection:      COLLECTION_NAME,
				Key:             object.Key,
				UserID:          params.UserId,
				Value:           string(data),
				PermissionRead:  1,
				PermissionWrite: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	}
	data, err := json.Marshal(
		params.Inventory,
	)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             uuid.NewString(),
			UserID:          params.UserId,
			Value:           string(data),
			PermissionRead:  1,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
