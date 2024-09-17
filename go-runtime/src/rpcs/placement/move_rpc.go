package rpcs_placement

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type MoveRpcParams struct {
	PlacedItemKey string                            `json:"placedItemKey"`
	Position      collections_placed_items.Position `json:"position"`
}

func MoveRpc(ctx context.Context,
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
	var params *MoveRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	object, err := collections_placed_items.ReadByKey(ctx, logger, db, nk, collections_placed_items.ReadByKeyParams{
		Key:    params.PlacedItemKey,
		UserId: userId,
	})
	if object == nil {
		errMsg := "placed item not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	placedItem, err := collections_common.ToValue[collections_placed_items.PlacedItem](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	placedItem.Position = params.Position
	_, err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: *placedItem,
		UserId:     userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return "", nil
}
