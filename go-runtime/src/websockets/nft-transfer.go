package websockets

import (
	services_periphery "cifarm-server/src/services/periphery"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
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

	client, err := socketio_client.NewClient(url, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	// Handle an incoming event
	client.On("connection", func() {
		logger.Info("Websocket to peripery server ebstablisted.")
	})
	client.On("nft-transfer-observed", func(msg string) {
		logger.Info(msg)
	})

	return nil
}
