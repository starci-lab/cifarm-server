package authenticator_api

import (
	_constants "cifarm-server/src/constants"
	_api "cifarm-server/src/utils/api"
	"context"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type VerifyMessageRequestBody struct {
	Message   string `json:"message"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
	Chain     string `json:"chain"`
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

func VerifyMessage(ctx context.Context, logger runtime.Logger, body *VerifyMessageRequestBody) (response *VerifyMessageResponseData, err error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[_constants.ENV_CI_BASE_API_URL]
	if !ok {
		logger.Error("CI_BASE_API_URL not found in environment variables")
		return nil, errors.New("CI_BASE_API_URL not found in environment variables")
	}
	url = url + "/authenticator/verify-message"
	logger.Info("POST %v", url)
	_response, err := _api.SendPostRequest[VerifyMessageRequestBody, VerifyMessageResponse](url, body)
	if err != nil {
		return nil, err
	}
	return &_response.Data, nil
}
