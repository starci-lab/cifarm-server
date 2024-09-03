package daily_reward

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdateWalletParams struct {
	UserId   string
	Amount   int64
	Metadata map[string]interface{}
}

func UpdateWallet(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params UpdateWalletParams,
) error {
	changeset := map[string]int64{
		_constants.KEY_GOLDS: params.Amount,
	}

	userId := params.UserId
	if userId == "" {
		_userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
		if !ok {
			errMsg := "user ID not found"
			logger.Error(errMsg)
			return errors.New(errMsg)
		}
		userId = _userId
	}

	_, _, err := nk.WalletUpdate(ctx, userId, changeset, params.Metadata, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
