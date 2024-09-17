package rpcs_tests

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func DeliveryRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	return "", nil
}
