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

type RecoverTileRpcParams struct {
	PlacedItemTileKey string                            `json:"inventoryTileKey"`
	Position          collections_placed_items.Position `json:"position"`
}

type RecoverTileRpcResponse struct {
	InventoryTileKey string `json:"inventoryTileKey"`
}

func RecoverTileRpc(ctx context.Context,
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
	var params *RecoverTileRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemTileKey,
		UserId: userId,
	})
	if object == nil {
		errMsg := "placed item not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	placedItem, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if placedItem.InventoryKey == "" {
		errMsg := "placed item inventory key not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	if placedItem.Type != collections_placed_items.TYPE_TILE {
		errMsg := "placed item not tile"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//delete the placed item
	err = collections_placed_items.Delete(ctx, logger, db, nk, collections_placed_items.DeleteParams{
		Key:    params.PlacedItemTileKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update inventory
	inventoryObject, err := collections_inventories.ReadByKey(ctx, logger, db, nk, collections_inventories.ReadByKeyParams{
		Key:    placedItem.InventoryKey,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	inventory, err := collections_common.ToValue[collections_inventories.Inventory](ctx, logger, db, nk, inventoryObject)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//update the tile
	inventory.IsPlaced = false
	result, err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
		Inventory: *inventory,
		UserId:    userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(RecoverTileRpcResponse{
		InventoryTileKey: result.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), nil
}
