package animals

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

type ReadAnimalObjectByIdParams struct {
	Id string `json:"id"`
}

func ReadAnimalObjectById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadAnimalObjectByIdParams,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_ANIMAL_OBJECTS
	query := fmt.Sprintf("+value.id:%s", params.Id)
	order := []string{}

	animals, err := nk.StorageIndexList(ctx, "", name, query, 100, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(animals.Objects) == 0 {
		errMsg := "animal not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	animal := animals.Objects[0]
	return animal, nil
}

func ReadAnimalObjectValueById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadAnimalObjectByIdParams,
) (*_collections.Animal, error) {
	object, err := ReadAnimalObjectById(ctx, logger, db, nk, params)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var animal *_collections.Animal
	err = json.Unmarshal([]byte(object.Value), &animal)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return animal, nil
}
