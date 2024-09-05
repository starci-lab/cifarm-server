package rpcs_shop

import (
	collections_buildings "cifarm-server/src/collections/buildings"
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ConstructBuildingRpcParams struct {
	Key      string                            `json:"key"`
	Position collections_placed_items.Position `json:"position"`
}

type ConstructBuildingRpcResponse struct {
	Cost int64 `json:"cost"`
}

func ConstructBuildingRpc(ctx context.Context,
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

	var params *ConstructBuildingRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	object, err := collections_buildings.ReadByKey(ctx, logger, db, nk, collections_buildings.ReadByKeyParams{
		Key: params.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	building, err := collections_common.ToValue[collections_buildings.Bulding](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if building == nil {
		errMsg := "building not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		UserId: userId,
		Amount: -building.Price,
		Metadata: map[string]interface{}{
			"name": "Construct building",
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	err = collections_placed_items.Write(ctx, logger, db, nk, collections_placed_items.WriteParams{
		PlacedItem: collections_placed_items.PlacedItem{
			ReferenceId: params.Key,
			Position:    params.Position,
			Type:        collections_placed_items.TYPE_BUILDING,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	value, err := json.Marshal(ConstructBuildingRpcResponse{
		Cost: building.Price,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), nil
}
