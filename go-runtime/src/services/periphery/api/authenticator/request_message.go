package services_periphery_api_authenticator

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type RequestMessageResponseData struct {
	Message string `json:"message"`
}

type RequestMessageResponse struct {
	Message string                     `json:"message"`
	Data    RequestMessageResponseData `json:"data"`
}

func RequestMessage(ctx context.Context, logger runtime.Logger) (response *RequestMessageResponseData, err error) {
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
	url = url + "/authenticator/request-message"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[any, RequestMessageResponse](url, nil)
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
