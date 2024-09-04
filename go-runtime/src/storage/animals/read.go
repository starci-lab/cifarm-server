package animals

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
	name := _constants.STORAGE_INDEX_ANIMALS
	query := fmt.Sprintf(`+value.id:%s`, params.Id)
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

func ToAnimal(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.Animal, error) {
	if object == nil {
		return nil, nil
	}
	var animal *_collections.Animal
	err := json.Unmarshal([]byte(object.Value), &animal)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return animal, nil
}
