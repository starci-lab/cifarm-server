package services_periphery_api_authenticator

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VerifyMessageRequestBody struct {
	Message   string `json:"message"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
	ChainKey  string `json:"chainKey"`
}

type VerifyMessageResponseData struct {
	Result  bool   `json:"result"`
	Address string `json:"address"`
}

type VerifyMessageResponse struct {
	Message string                    `json:"message"`
	Data    VerifyMessageResponseData `json:"data"`
}

func VerifyMessage(ctx context.Context, logger runtime.Logger, body *VerifyMessageRequestBody) (response *VerifyMessageResponseData, err error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[cifarm_periphery.CIFARM_PERIPHERY_API_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_API_URL not found in environment variables")
		return nil, errors.New("CIFARM_PERIPHERY_API_URL not found in environment variables")
	}
	url = url + "/authenticator/verify-message"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[VerifyMessageRequestBody, VerifyMessageResponse](url, body)
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
