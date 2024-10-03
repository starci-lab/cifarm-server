package rpcs_tests

import (
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type GetDeliveringProductsRpcResponse struct {
	DeliveringProductBasicKey   string `json:"deliveringProductBasicKey"`
	DeliveringProductPremiumKey string `json:"deliveringProductPremiumKey"`
}

func GetDeliveringProductsRpc(ctx context.Context,
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

	//this will send you 20 carrot and 20 premium carrot
	basicResult, err := collections_delivering_products.Write(ctx, logger, db, nk, collections_delivering_products.WriteParams{
		DeliveringProduct: collections_delivering_products.DeliveringProduct{
			ReferenceKey: "carrot",
			Type:         collections_delivering_products.TYPE_CROP,
			Quantity:     20,
			Index:        1,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	premiumResult, err := collections_delivering_products.Write(ctx, logger, db, nk, collections_delivering_products.WriteParams{
		DeliveringProduct: collections_delivering_products.DeliveringProduct{
			ReferenceKey: "carrot",
			Type:         collections_delivering_products.TYPE_CROP,
			Quantity:     20,
			Index:        2,
			Premium:      true,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(GetDeliveringProductsRpcResponse{
		DeliveringProductBasicKey:   basicResult.Key,
		DeliveringProductPremiumKey: premiumResult.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(value), nil
}
