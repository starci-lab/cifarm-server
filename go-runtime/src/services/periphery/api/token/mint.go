package services_periphery_api_token

import (
	"cifarm-server/src/config"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type MintRequestBody struct {
	TokenAddress     string  `json:"tokenAddress"`
	ToAddress        string  `json:"toAddress"`
	MinterPrivateKey string  `json:"minterPrivateKey"`
	MintAmount       float64 `json:"mintAmount"`
	ChainKey         string  `json:"chainKey"`
	Network          string  `json:"network"`
}

type MintResponseData struct {
	TransactionHash string `json:"transactionHash"`
}

type MintResponse struct {
	Message string           `json:"message"`
	Data    MintResponseData `json:"data"`
}

func Mint(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, body *MintRequestBody) (response *MintResponseData, err error) {
	url, err := config.CifarmPeripheryApiUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url = url + "/token/mint"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[MintRequestBody, MintResponse](url, body, nil)
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
