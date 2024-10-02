package collections_delivering_products

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	DeliveringProduct DeliveringProduct `json:"deliveringProduct"`
	UserId            string            `json:"userId"`
}

type WriteResult struct {
	Key string `json:"key"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) (*WriteResult, error) {
	object, err := Read(ctx, logger, db, nk, ReadParams{
		ReferenceKey: params.DeliveringProduct.ReferenceKey,
		UserId:       params.UserId,
		Index:        params.DeliveringProduct.Index,
		Premium:      params.DeliveringProduct.Premium,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if object != nil {
		deliveringProduct, err := collections_common.ToValue[DeliveringProduct](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		deliveringProduct.Quantity += params.DeliveringProduct.Quantity
		data, err := json.Marshal(deliveringProduct)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		acks, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{
			{
				Collection:      COLLECTION_NAME,
				Key:             object.Key,
				UserID:          params.UserId,
				Value:           string(data),
				PermissionRead:  2,
				PermissionWrite: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		result := &WriteResult{
			Key: acks[0].Key,
		}
		return result, nil
	}

	key := uuid.NewString()

	data, err := json.Marshal(
		params.DeliveringProduct,
	)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	acks, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
			Value:           string(data),
			PermissionRead:  1,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	result := &WriteResult{
		Key: acks[0].Key,
	}
	return result, nil
}
