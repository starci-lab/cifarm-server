package services_periphery_api_authenticator

import (
	"cifarm-server/src/config"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type RequestMessageResponseData struct {
	Message string `json:"message"`
}

type RequestMessageResponse struct {
	Message string                     `json:"message"`
	Data    RequestMessageResponseData `json:"data"`
}

func RequestMessage(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (response *RequestMessageResponseData, err error) {
	url, err := config.CifarmPeripheryApiUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url = url + "/authenticator/request-message"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[any, RequestMessageResponse](url, nil)
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
