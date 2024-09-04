package shop

import (
	_constants "cifarm-server/src/constants"
	_farming_tiles "cifarm-server/src/storage/farming_tiles"
	_inventories "cifarm-server/src/storage/inventories"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HasFarmingTileParams struct {
	Id string `json:"id"`
}

func HasEnoughFarmingTiles(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params HasFarmingTileParams) (bool, error) {
	object, err := _inventories.ReadInventoryObject(ctx, logger, db, nk, _inventories.ReadInventoryObjectParams{
		Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_1,
	},
	)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	inventory, err := _inventories.ToInventory(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	object, err = _farming_tiles.ReadFarmingTileObjectById(
		ctx, logger, db, nk,
		_farming_tiles.ReadFarmingTileObjectByIdParams{
			Id: params.Id,
		})
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	farmingTile, err := _farming_tiles.ToFarmingTile(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}
	result := inventory.Quantity >= farmingTile.MaxOwnership
	return result, nil
}

type GetFarmingTileIdAndPriceResult struct {
	Id    string `json:"id"`
	Price int64  `json:"price"`
}

func GetFarmingTileIdAndPrice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (*GetFarmingTileIdAndPriceResult, error) {
	has1, err := HasEnoughFarmingTiles(ctx, logger, db, nk, HasFarmingTileParams{
		Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_1,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if !has1 {
		object, err := _farming_tiles.ReadFarmingTileObjectById(
			ctx, logger, db, nk,
			_farming_tiles.ReadFarmingTileObjectByIdParams{
				Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_1,
			})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		farmingTile, err := _farming_tiles.ToFarmingTile(ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		id := _constants.FARMING_TILE_BASIC_FARMING_TILE_1
		price := farmingTile.Price

		return &GetFarmingTileIdAndPriceResult{
			Id:    id,
			Price: price,
		}, nil
	}
	has2, err := HasEnoughFarmingTiles(ctx, logger, db, nk, HasFarmingTileParams{
		Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_2,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if !has2 {
		object, err := _farming_tiles.ReadFarmingTileObjectById(
			ctx, logger, db, nk,
			_farming_tiles.ReadFarmingTileObjectByIdParams{
				Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_2,
			})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		farmingTile, err := _farming_tiles.ToFarmingTile(ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		id := _constants.FARMING_TILE_BASIC_FARMING_TILE_2
		price := farmingTile.Price

		return &GetFarmingTileIdAndPriceResult{
			Id:    id,
			Price: price,
		}, nil
	}
	object, err := _farming_tiles.ReadFarmingTileObjectById(
		ctx, logger, db, nk,
		_farming_tiles.ReadFarmingTileObjectByIdParams{
			Id: _constants.FARMING_TILE_BASIC_FARMING_TILE_3,
		})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	farmingTile, err := _farming_tiles.ToFarmingTile(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	id := _constants.FARMING_TILE_BASIC_FARMING_TILE_3
	price := farmingTile.Price

	return &GetFarmingTileIdAndPriceResult{
		Id:    id,
		Price: price,
	}, nil
}

type BuyFarmingTileRpcResponse struct {
	Id    string `json:"id"`
	Price int64  `json:"price"`
}

func BuyFarmingTileRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	data, err := GetFarmingTileIdAndPrice(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		UserId: userId,
		Amount: -data.Price,
		Metadata: map[string]interface{}{
			"name":   "Buy farming tile",
			"seedId": data.Id,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	err = _inventories.WriteInventoryObject(ctx,
		logger, db, nk,
		_inventories.WriteInventoryObjectParams{
			Id:       data.Id,
			Quantity: 1,
			Type:     _constants.TYPE_FARMING_TILE,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}