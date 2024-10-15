package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_tools "cifarm-server/src/collections/tools"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

// player's tools
type PlayerTool struct {
	Key           string                            `json:"key"`
	FromInventory bool                              `json:"fromInventory"`
	Inventory     collections_inventories.Inventory `json:"inventory"`
	Type          int                               `json:"type"`
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

	// query tools
	// list all defaut tools
	defaultToolObjects, err := collections_tools.ReadMany(ctx, logger, db, nk, collections_tools.ReadManyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	defaultTools, err := collections_common.ToValues2[collections_tools.Tool](ctx, logger, db, nk, defaultToolObjects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	var tools []PlayerTool
	for _, defaultTool := range defaultTools {
		tools = append(tools, PlayerTool{
			Key:           defaultTool.Key,
			FromInventory: false,
		})
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
			Inventory:     *inventory,
		})
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
