package rpcs_assets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ListDeliveringProductsRpcResponse struct {
	DeliveringProducts []*collections_delivering_products.DeliveringProduct `json:"deliveringProducts"`
}

func ListDeliveringProductsRpc(
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

	objects, err := collections_delivering_products.ReadMany(ctx, logger, db, nk, collections_delivering_products.ReadManyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	deliveringProducts, err := collections_common.ToValues2[collections_delivering_products.DeliveringProduct](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(ListDeliveringProductsRpcResponse{
		DeliveringProducts: deliveringProducts,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), err
}
