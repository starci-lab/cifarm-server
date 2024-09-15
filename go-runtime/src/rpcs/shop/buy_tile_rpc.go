package rpcs_shop

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HasEnoughTilesParams struct {
	ReferenceKey string `json:"referenceKey"`
}

func HasEnoughTiles(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params HasEnoughTilesParams,
) (bool, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return false, errors.New(errMsg)
	}

	object, err := collections_inventories.ReadByReferenceKey(ctx, logger, db, nk, collections_inventories.ReadByReferenceKeyParams{
		ReferenceKey: params.ReferenceKey,
		UserId:       userId,
	},
	)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	inventory, err := collections_common.ToValue[collections_inventories.Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	if inventory == nil {
		return false, nil
	}

	object, err = collections_tiles.ReadByKey(ctx, logger, db, nk, collections_tiles.ReadByKeyParams{
		Key: params.ReferenceKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	tile, err := collections_common.ToValue[collections_tiles.Tile](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	if tile == nil {
		errMsg := "tile not found"
		logger.Error(errMsg)
		return false, errors.New(errMsg)
	}

	result := inventory.Quantity >= tile.MaxOwnership
	return result, nil
}

type GetTileDataResult struct {
	Key   string `json:"key"`
	Price int64  `json:"price"`
}

func GetTileData(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*GetTileDataResult, error) {
	key := collections_tiles.KEY_BASIC_1
	has1, err := HasEnoughTiles(ctx, logger, db, nk, HasEnoughTilesParams{
		ReferenceKey: key,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if !has1 {
		object, err := collections_tiles.ReadByKey(ctx, logger, db, nk, collections_tiles.ReadByKeyParams{
			Key: key,
		})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		tile, err := collections_common.ToValue[collections_tiles.Tile](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		if tile == nil {
			errMsg := "tile not found"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}

		price := tile.Price

		return &GetTileDataResult{
			Key:   key,
			Price: price,
		}, nil
	}
	key = collections_tiles.KEY_BASIC_2
	has2, err := HasEnoughTiles(ctx, logger, db, nk, HasEnoughTilesParams{
		ReferenceKey: key,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if !has2 {
		object, err := collections_tiles.ReadByKey(ctx, logger, db, nk, collections_tiles.ReadByKeyParams{
			Key: key,
		})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		tile, err := collections_common.ToValue[collections_tiles.Tile](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}

		if tile == nil {
			errMsg := "tile not found"
			logger.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		price := tile.Price

		return &GetTileDataResult{
			Key:   key,
			Price: price,
		}, nil
	}
	key = collections_tiles.KEY_BASIC_3
	object, err := collections_tiles.ReadByKey(ctx, logger, db, nk, collections_tiles.ReadByKeyParams{
		Key: key,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	tile, err := collections_common.ToValue[collections_tiles.Tile](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	price := tile.Price

	return &GetTileDataResult{
		Key:   key,
		Price: price,
	}, nil
}

type BuyTileRpcResponse struct {
	Key   string `json:"key"`
	Price int64  `json:"price"`
}

func BuyTileRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	data, err := GetTileData(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		Amount: -data.Price,
		Metadata: map[string]interface{}{
			"name": "Buy tile",
			"key":  data.Key,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	_, err = collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: collections_inventories.Inventory{
			ReferenceKey: data.Key,
			Quantity:     1,
			Type:         collections_inventories.TYPE_TILE,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuyTileRpcResponse{
		Price: data.Price,
		Key:   data.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
