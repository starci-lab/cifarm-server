package rpcs_tests

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type SpeedUpRpcParams struct {
	Time int64 `json:"time"`
}

func SpeedUpRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	var params *SpeedUpRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = collections_system.WriteSpeedUp(ctx, logger, db, nk, collections_system.WriteSpeedUpParams{
		SpeedUp: collections_system.SpeedUp{
			AnimalProcedureTime: params.Time,
			SeedGrowthTime:      params.Time,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
