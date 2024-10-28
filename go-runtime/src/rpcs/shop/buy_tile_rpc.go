package rpcs_shop

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HasEnoughTilesParams struct {
	ReferenceKey string `json:"referenceKey,omitempty"`
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

	objects, err := collections_placed_items.ReadByFilters3(ctx, logger, db, nk, collections_placed_items.ReadByFilters3Params{
		UserId:       userId,
		ReferenceKey: params.ReferenceKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	tiles, err := collections_common.ToValues[collections_placed_items.PlacedItem](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	object, err := collections_tiles.ReadByKey(ctx, logger, db, nk, collections_tiles.ReadByKeyParams{
		Key: params.ReferenceKey,
	})
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	if object == nil {
		errMsg := "tile not found"
		logger.Error(errMsg)
		return false, errors.New(errMsg)
	}
	tile, err := collections_common.ToValue[collections_tiles.Tile](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	result := len(tiles) >= tile.MaxOwnership
	return result, nil
}

type GetTileDataResult struct {
	Key   string `json:"key,omitempty"`
	Price int64  `json:"price,omitempty"`
}

func GetTileData(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*GetTileDataResult, error) {
	key := collections_tiles.KEY_BASIC_TILE_1
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
	key = collections_tiles.KEY_BASIC_TILE_2
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
	key = collections_tiles.KEY_BASIC_TILE_3
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

type BuyTileRpcParams struct {
	Position collections_placed_items.Position `json:"position,omitempty"`
}

type BuyTileRpcResponse struct {
	PlacedItemTileKey string `json:"placedItemTileKey,omitempty"`
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

	var params *BuyTileRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	data, err := GetTileData(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		GoldAmount: -data.Price,
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
	result, err := collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: collections_placed_items.PlacedItem{
			ReferenceKey: data.Key,
			Type:         collections_placed_items.TYPE_TILE,
			Position:     params.Position,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(BuyTileRpcResponse{
		PlacedItemTileKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
