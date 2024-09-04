package farming

import (
	_constants "cifarm-server/src/constants"
	_inventories "cifarm-server/src/storage/inventories"
	_placed_items "cifarm-server/src/storage/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type PlantSeedRpcParams struct {
	InventoryKey         string `json:"inventoryKey"`
	PlacedFarmingTileKey string `json:"placedFarmingTileKey"`
}

func PlantSeedRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	var params *PlantSeedRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := _inventories.ReadInventoryObjectByKey(ctx, logger, db, nk, params.InventoryKey)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	inventory, err := _inventories.ToInventory(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if inventory == nil {
		errMsg := "inventory not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if inventory.Type != _constants.INVENTORY_TYPE_PLANT_SEED {
		errMsg := "inventory not plant seed"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	object, err = _placed_items.ReadPlacedItemObjectByKey(ctx, logger, db, nk, params.PlacedFarmingTileKey)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	placedItem, err := _placed_items.ToPlacedItem(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if placedItem.Type != _constants.PLACED_ITEM_TYPE_FARMING_TILE {
		errMsg := "placed item not farming tile"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	return "", nil
}
