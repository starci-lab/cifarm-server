package rpcs_tests

import (
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HackGoldRpcParams struct {
	Amount int64 `json:"amount"`
}

func HackGoldRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *HackGoldRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		GoldAmount:  params.Amount,
		TokenAmount: float64(params.Amount),
		UserId:      userId,
		Metadata: map[string]interface{}{
			"name": "Hack",
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
