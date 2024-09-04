package plant_seeds

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func ReadPlantSeedObjectById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	id string,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_PLANT_SEEDS
	query := fmt.Sprintf(`+value.id:%s`, id)
	order := []string{}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 100, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}

	object := objects.Objects[0]
	return object, nil
}

func ToPlantSeed(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.PlantSeed, error) {
	if object == nil {
		return nil, nil
	}
	var plantSeed *_collections.PlantSeed
	err := json.Unmarshal([]byte(object.Value), &plantSeed)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return plantSeed, nil
}
