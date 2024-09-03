package plant_seeds

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadPlantSeedObjectByIdParams struct {
	Id string `json:"Id"`
}

func ReadPlantSeedObjectById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadPlantSeedObjectByIdParams,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_PLANT_SEED_OBJECTS
	query := fmt.Sprintf("+value.id:%s", params.Id)
	order := []string{}

	plantSeeds, err := nk.StorageIndexList(ctx, "", name, query, 100, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(plantSeeds.Objects) == 0 {
		errMsg := "plant seed not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	plantSeed := plantSeeds.Objects[0]
	return plantSeed, nil
}

func ReadPlantSeedObjectValueById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadPlantSeedObjectByIdParams,
) (*_collections.PlantSeed, error) {
	object, err := ReadPlantSeedObjectById(ctx, logger, db, nk, params)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var plantSeed *_collections.PlantSeed
	err = json.Unmarshal([]byte(object.Value), &plantSeed)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return plantSeed, nil
}
