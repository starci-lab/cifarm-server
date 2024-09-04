package shop

import (
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

type BuyLandRpcParams struct {
	Quantity int `json:"quantity"`
}

type BuyLandRpcResponse struct {
	TotalCost int64 `json:"totalCost"`
}

func BuyLandRpc(ctx context.Context,
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
	var params *BuyPlantSeedRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := _plant_seeds.ReadPlantSeedObjectById(
		ctx, logger, db, nk,
		_plant_seeds.ReadPlantSeedObjectByIdParams{
			Id: params.Id,
		})
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

	totalCost := int64(plantSeed.SeedPrice) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		UserId: userId,
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
		_inventories.WriteInventoryObjectParams{
			Id:       plantSeed.Id,
			Quantity: params.Quantity,
			Type:     _collections.TYPE_SEED,
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
