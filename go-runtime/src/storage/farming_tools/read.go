package farming_tools

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

type ReadFarmingToolObjectByIdParams struct {
	Id string `json:"Id"`
}

func ReadFarmingToolObjectById(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadFarmingToolObjectByIdParams,
) (*api.StorageObject, error) {
	name := _constants.STORAGE_INDEX_FARMING_TOOLS
	query := fmt.Sprintf("+value.id:%s", params.Id)
	order := []string{
		"-create_time",
	}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 100, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}
	var object = objects.Objects[0]
	return object, nil
}

func ToFarmingTool(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.FarmingTool, error) {
	var farmingTool *_collections.FarmingTool
	if object == nil {
		return nil, nil
	}
	err := json.Unmarshal([]byte(object.Value), &farmingTool)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return farmingTool, nil
}
