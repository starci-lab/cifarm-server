package websockets

import (
	services_periphery "cifarm-server/src/services/periphery"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

func ObserveNftTransfer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return errors.New("cannot get environment variables")
	}
	url, ok := vars[services_periphery.CIFARM_PERIPHERY_WEBSOCKET_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_WEBSOCKET_URL not found in environment variables")
		return errors.New("CIFARM_PERIPHERY_WEBSOCKET_URL not found in environment variables")
	}
	logger.Error(url)
	return nil
}
