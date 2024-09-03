package shop

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	_wallets "cifarm-server/src/utils/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

	name := _constants.STORAGE_INDEX_PLANT_SEED_OBJECTS
	query := fmt.Sprintf("value.id:%s", params.Id)
	order := []string{}

	plantSeeds, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if len(plantSeeds.Objects) == 0 {
		errMsg := "plant seed not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var plantSeed = plantSeeds.Objects[0]

	var _plantSeed *_collections.PlantSeed
	err = json.Unmarshal([]byte(plantSeed.Value), &_plantSeed)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	totalCost := int64(_plantSeed.SeedPrice) * int64(params.Quantity)
	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		UserId: userId,
		Amount: -totalCost,
		Metadata: map[string]interface{}{
			"name":   "Buy seeds",
			"seedId": _plantSeed.Id,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	inventory, err := json.Marshal(_collections.Inventory{
		Id:       _plantSeed.Id,
		Type:     _collections.TYPE_SEED,
		Quantity: params.Quantity,
	})

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      _constants.COLLECTION_INVENTORY,
			Key:             uuid.NewString(),
			UserID:          userId,
			Value:           string(inventory),
			PermissionRead:  1,
			PermissionWrite: 0,
		},
	})

	_value, err := json.Marshal(BuyPlantSeedRpcResponse{
		TotalCost: totalCost,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
