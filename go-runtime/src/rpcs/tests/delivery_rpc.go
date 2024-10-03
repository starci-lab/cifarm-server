package rpcs_tests

import (
	crons_deliver "cifarm-server/src/crons/deliver"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func DeliveryRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	//do the same logic as process
	err := crons_deliver.Process(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
