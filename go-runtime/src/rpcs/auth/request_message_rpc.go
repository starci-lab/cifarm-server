package rpcs_auth

import (
	services_periphery_api_authenticator "cifarm-server/src/services/periphery/api/authenticator"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type RequestMessageRpcResponse struct {
	Message string `json:"message"`
}

func RequestMessageRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	response, err := services_periphery_api_authenticator.RequestMessage(ctx, logger, db, nk)
	if err != nil {
		return "", err
	}

	_response := &RequestMessageRpcResponse{Message: response.Message}

	out, err := json.Marshal(_response)
	if err != nil {
		return "", nil
	}

	return string(out), nil
}
