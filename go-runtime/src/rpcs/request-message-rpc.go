package rpcs

import (
	_authenticator "cifarm-server/src/services/ci-base/authenticator"
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
	response, err := _authenticator.RequestMessage(ctx, logger)
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
