package shop

import (
	_constants "cifarm-server/src/constants"
	_inventories "cifarm-server/src/storage/inventories"
	_plant_seeds "cifarm-server/src/storage/plant_seeds"
	_collections "cifarm-server/src/types/collections"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuyPlantSeedRpcParams struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type BuyPlantSeedRpcResponse struct {
	TotalCost int64 `json:"totalCost"`
}

func BuyPlantSeedRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	var params *BuyPlantSeedRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := _plant_seeds.ReadPlantSeedObjectById(ctx, logger, db, nk, params.Id)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	plantSeed, err := _plant_seeds.ToPlantSeed(
		ctx,
		logger,
		db,
		nk,
		object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if plantSeed == nil {
		errMsg := "plant seed not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	totalCost := int64(plantSeed.SeedPrice) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		Amount: -totalCost,
		Metadata: map[string]interface{}{
			"name":   "Buy seeds",
			"seedId": plantSeed.Id,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = _inventories.WriteInventoryObject(ctx,
		logger, db, nk,
		_collections.Inventory{
			Id:       plantSeed.Id,
			Quantity: params.Quantity,
			Type:     _constants.INVENTORY_TYPE_PLANT_SEED,
		})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	_value, err := json.Marshal(BuyPlantSeedRpcResponse{
		TotalCost: totalCost,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
