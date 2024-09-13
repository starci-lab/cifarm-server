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
	object, err := ReadByReferenceKey(ctx, logger, db, nk, ReadByReferenceKeyParams{
		ReferenceKey: params.Inventory.ReferenceKey,
		UserId:       params.UserId,
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
		inventory.Quantity += params.Inventory.Quantity
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

	key := uuid.NewString()
	params.Inventory.Key = key

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
			Key:             key,
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

type WriteUniqueParams struct {
	Inventory      Inventory `json:"inventory"`
	UserId         string    `json:"userId"`
	PermissionRead int       `json:"permissionRead"`
}

func WriteUnique(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteUniqueParams,
) error {
	key := uuid.NewString()
	params.Inventory.Key = key
	params.Inventory.Unique = true

	value, err := json.Marshal(params.Inventory)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
			Value:           string(value),
			PermissionRead:  params.PermissionRead,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type TransferOwnershipParams struct {
	Key        string `json:"key"`
	FromUserId string `json:"fromUserId"`
	ToUserId   string `json:"toUserId"`
}

func TransferOwnership(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params TransferOwnershipParams,
) error {
	object, err := ReadByKey(ctx, logger, db, nk, ReadByKeyParams{
		Key:    params.Key,
		UserId: params.FromUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	value, err := json.Marshal(object.Value)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	//we do destroy placed items

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             object.Key,
			UserID:          params.ToUserId,
			Value:           string(value),
			PermissionRead:  int(object.PermissionRead),
			PermissionWrite: int(object.PermissionWrite),
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
