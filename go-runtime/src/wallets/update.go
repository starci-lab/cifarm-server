package wallets

import (
	"context"
	"database/sql"

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
		WALLETS_KEY_GOLD: params.Amount,
	}

	_, _, err := nk.WalletUpdate(ctx, params.UserId, changeset, params.Metadata, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
