package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_player "cifarm-server/src/collections/player"
	collections_tools "cifarm-server/src/collections/tools"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"sort"

	"github.com/heroiclabs/nakama-common/runtime"
)

// player's tools
type PlayerTool struct {
	Key           string `json:"key"`
	FromInventory bool   `json:"fromInventory"`
	InventoryKey  string `json:"inventoryKey"`
	Type          int    `json:"type"`
	Index         int    `json:"index"`
}

type ListToolsRpcResponse struct {
	Tools []PlayerTool `json:"tools"`
}

func ListToolsRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//load the visit
	object, err := collections_player.ReadVisitState(ctx, logger, db, nk, collections_player.ReadVisitStateParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "visit state not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	visitState, err := collections_common.ToValue[collections_player.VisitState](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//visit other
	visited := visitState.UserId != ""

	// query tools
	// list all defaut tools
	defaultToolObjects, err := collections_tools.ReadMany(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	defaultTools, err := collections_common.ToValues2[collections_tools.Tool](ctx, logger, db, nk, defaultToolObjects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	sort.Slice(defaultTools, func(i, j int) bool {
		return defaultTools[i].Index < defaultTools[j].Index
	})

	var index int
	var tools []PlayerTool
	for _, defaultTool := range defaultTools {
		if !(defaultTool.AvailableIn == collections_tools.AVAILABLE_IN_NEIGHBOR && !visited ||
			defaultTool.AvailableIn == collections_tools.AVAILABLE_IN_HOME && visited) {
			tools = append(tools, PlayerTool{
				Key:           defaultTool.Key,
				FromInventory: false,
				Index:         index,
			})
			index++
		}
	}

	//load tools from inventory
	inventoryObjects, err := collections_inventories.ReadMany(ctx, logger, db, nk, collections_inventories.ReadManyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	inventories, err := collections_common.ToValues2[collections_inventories.Inventory](ctx, logger, db, nk, inventoryObjects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	for _, inventory := range inventories {
		if !inventory.AsTool {
			continue
		}
		tools = append(tools, PlayerTool{
			Key:           inventory.Key,
			FromInventory: true,
			InventoryKey:  inventory.Key,
			Index:         index,
		})
		index++
	}

	value, err := json.Marshal(ListToolsRpcResponse{
		Tools: tools,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
