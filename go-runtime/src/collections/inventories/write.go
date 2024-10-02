package collections_inventories

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	Inventory Inventory `json:"inventory"`
	UserId    string    `json:"userId"`
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
	//read only non delivering
	object, err := Read(ctx, logger, db, nk, ReadParams{
		ReferenceKey: params.Inventory.ReferenceKey,
		UserId:       params.UserId,
		Type:         params.Inventory.Type,
		Premium:      params.Inventory.Premium,
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if object != nil {
		inventory, err := collections_common.ToValue[Inventory](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		inventory.Quantity += params.Inventory.Quantity
		data, err := json.Marshal(inventory)
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
		params.Inventory,
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

type WriteUniqueParams struct {
	Inventory Inventory `json:"inventory"`
	UserId    string    `json:"userId"`
}

type WriteUniqueResult struct {
	Key string `json:"key"`
}

func WriteUnique(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteUniqueParams,
) (*WriteUniqueResult, error) {
	key := params.Inventory.Key
	if key == "" {
		key = uuid.NewString()
	}
	params.Inventory.Unique = true

	value, err := json.Marshal(params.Inventory)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &WriteUniqueResult{Key: key}, nil
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
	if object == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	err = DeleteUnique(ctx, logger, db, nk, DeleteUniqueParams{
		Key:    object.Key,
		UserId: params.FromUserId,
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             object.Key,
			UserID:          params.ToUserId,
			Value:           object.Value,
			PermissionRead:  int(object.PermissionRead),
			PermissionWrite: int(object.PermissionWrite),
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_placed_items.DeleteByInventoryKey(ctx, logger, db, nk, collections_placed_items.DeleteByInventoryKeyParams{
		InventoryKey: params.Key,
		UserId:       params.FromUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

type UpdateIsPlacedParams struct {
	Key      string `json:"key"`
	IsPlaced bool   `json:"isPlaced"`
	UserId   string `json:"userId"`
}

func UpdateIsPlaced(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params UpdateIsPlacedParams,
) error {
	object, err := ReadByKey(ctx, logger, db, nk, ReadByKeyParams{
		Key:    params.Key,
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	inventory, err := collections_common.ToValue[Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	inventory.IsPlaced = params.IsPlaced
	value, err := json.Marshal(inventory)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             object.Key,
			UserID:          params.UserId,
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
