package blockchain_auth_service

import (
	_constants "cifarm-server/src/constants"
	_api "cifarm-server/src/utils/api"
	"context"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VerifyMessageRequestBody struct {
	Message   string  `json:"message"`
	Signature string  `json:"signature"`
	PublicKey string  `json:"public_key"`
	Platform  *string `json:"platform,omitempty"`
}

type VerifyMessageResponseData struct {
	Result  bool   `json:"result"`
	Address string `json:"address"`
}

func VerifyMessage(ctx context.Context, logger runtime.Logger, body *VerifyMessageRequestBody) (*VerifyMessageResponseData, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[_constants.ENV_BLOCKCHAIN_AUTH_SERVER_URL]
	if !ok {
		logger.Error("ENV_BLOCKCHAIN_AUTH_SERVER_URL not found in environment variables")
		return nil, errors.New("ENV_BLOCKCHAIN_AUTH_SERVER_URL not found in environment variables")
	}
	response, err := _api.SendPostRequest[VerifyMessageRequestBody, VerifyMessageResponseData](url, body)
	if err != nil {
		return nil, err
	}
	return response, nil
}
