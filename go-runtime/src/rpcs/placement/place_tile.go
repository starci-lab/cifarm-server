package rpcs_placement

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type PlaceTileRpcParams struct {
	InventoryTileKey string                            `json:"inventoryTileKey"`
	Position         collections_placed_items.Position `json:"position"`
}

type PlaceTileRpcResponse struct {
	PlacedItemTileKey string `json:"placedItemTileKey"`
}

func PlaceTileRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	var params *PlaceTileRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	object, err := collections_inventories.ReadByKey(ctx, logger, db, nk, collections_inventories.ReadByKeyParams{
		Key:    params.InventoryTileKey,
		UserId: userId,
	})
	if object == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	inventory, err := collections_common.ToValue[collections_inventories.Inventory](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if inventory.Type != collections_inventories.TYPE_TILE {
		errMsg := "inventory not tile"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if inventory.IsPlaced {
		errMsg := "inventory has been placed"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	result, err := collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: collections_placed_items.PlacedItem{
			ReferenceKey: inventory.ReferenceKey,
			InventoryKey: inventory.Key,
			Position:     params.Position,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update the tile
	inventory.IsPlaced = true
	_, err = collections_inventories.WriteUnique(ctx, logger, db, nk, collections_inventories.WriteUniqueParams{
		Inventory: *inventory,
		UserId:    userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(PlaceTileRpcResponse{
		PlacedItemTileKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), nil
}
