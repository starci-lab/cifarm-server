package services_periphery_api_authenticator

import (
	"cifarm-server/src/config"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VerifyMessageRequestBody struct {
	Message             string `json:"message,omitempty"`
	PublicKey           string `json:"publicKey,omitempty"`
	Signature           string `json:"signature,omitempty"`
	ChainKey            string `json:"chainKey,omitempty"`
	Network             string `json:"network,omitempty"`
	TelegramInitDataRaw string `json:"telegramInitDataRaw,omitempty"`
	BotType             string `json:"botType,omitempty"`
}

type VerifyMessageResponseData struct {
	Result           bool   `json:"result,omitempty"`
	AuthenticationId string `json:"authenticationId,omitempty"`
}

type VerifyMessageParams struct {
	Body VerifyMessageRequestBody `json:"body,omitempty"`
}

type VerifyMessageResponse struct {
	Message string                    `json:"message,omitempty"`
	Data    VerifyMessageResponseData `json:"data,omitempty"`
}

func VerifyMessage(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params VerifyMessageParams) (response *VerifyMessageResponseData, err error) {
	url, err := config.CifarmPeripheryApiUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url = url + "/authenticator/verify-message"

	body := params.Body
	_response, err := services_uitls_api.SendPostRequest[VerifyMessageRequestBody, VerifyMessageResponse](url, &body, nil)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &_response.Data, nil
}
