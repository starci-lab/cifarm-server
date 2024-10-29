package services_periphery_api_authenticator

import (
	"cifarm-server/src/config"
	services_uitls_api "cifarm-server/src/services/utils/api"
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/runtime"
)

type TelegramData struct {
	UserId int `json:"userId,omitempty"`
}

type AuthorizeTelegramResponseData struct {
	TelegramData TelegramData `json:"telegramData,omitempty"`
}

type AuthorizeTelegramResponse struct {
	Message string                        `json:"message,omitempty"`
	Data    AuthorizeTelegramResponseData `json:"data,omitempty"`
}

type AuthorizeTelegramParams struct {
	TelegramInitDataRaw string `json:"telegramInitDataRaw,omitempty"`
	BotType             string `json:"botType,omitempty"`
}

func AuthorizeTelegram(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params AuthorizeTelegramParams,
) (response *AuthorizeTelegramResponseData, err error) {
	url, err := config.CifarmPeripheryApiUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url = url + "/authenticator/authorize-telegram"
	logger.Info("POST %v", url)
	_response, err := services_uitls_api.SendPostRequest[any, AuthorizeTelegramResponse](url, nil, &map[string]string{
		"Authorization": fmt.Sprintf("tma %s", params.TelegramInitDataRaw),
		"Bot-Type":      params.BotType,
	})
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
