package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ListInventoriesRpcResponse struct {
	Inventories []*collections_inventories.Inventory `json:"inventories"`
}

func ListInventoriesRpc(
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

	objects, err := collections_inventories.ReadManyAvailable(ctx, logger, db, nk, collections_inventories.ReadManyAvailableParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	inventories, err := collections_common.ToValues[collections_inventories.Inventory](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(ListInventoriesRpcResponse{
		Inventories: inventories,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
