package services_periphery_api_authenticator

import (
	"cifarm-server/src/config"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VerifyMessageRequestBody struct {
	Message   string `json:"message"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
	ChainKey  string `json:"chainKey"`
	Network   string `json:"network"`
}

type VerifyMessageResponseData struct {
	Result           bool   `json:"result"`
	Address          string `json:"address"`
	AuthenticationId string `json:"authenticationId"`
}

type VerifyMessageResponse struct {
	Message string                    `json:"message"`
	Data    VerifyMessageResponseData `json:"data"`
}

func VerifyMessage(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, body *VerifyMessageRequestBody) (response *VerifyMessageResponseData, err error) {
	url, err := config.CifarmPeripheryApiUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url = url + "/authenticator/verify-message"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[VerifyMessageRequestBody, VerifyMessageResponse](url, body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &_response.Data, nil
}
