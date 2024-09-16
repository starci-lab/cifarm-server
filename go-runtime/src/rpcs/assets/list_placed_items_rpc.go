package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ListPlacedItemsRpcResponse struct {
	PlacedItems []*collections_placed_items.PlacedItem `json:"placedItems"`
}

func ListPlacedItemsRpc(
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

	objects, err := collections_placed_items.ReadMany(ctx, logger, db, nk, collections_placed_items.ReadManyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	placedItems, err := collections_common.ToValues2[collections_placed_items.PlacedItem](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(ListPlacedItemsRpcResponse{
		PlacedItems: placedItems,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
