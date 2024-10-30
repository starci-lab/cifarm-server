package wallets

import (
	collections_common "cifarm-server/src/collections/common"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"math"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdateWalletParams struct {
	UserId      string                 `json:"user_id"`
	GoldAmount  int64                  `json:"goldAmount"`
	TokenAmount float64                `json:"tokenAmount"`
	Metadata    map[string]interface{} `json:"metadata"`
}

func UpdateWallet(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params UpdateWalletParams,
) error {
	var goldUpdate, tokenUpdate int64
	if params.GoldAmount != 0 {
		goldUpdate = params.GoldAmount
	}
	if params.TokenAmount != 0 {
		object, err := collections_system.ReadTokenConfigure(ctx, logger, db, nk)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		tokenConfigure, err := collections_common.ToValue[collections_system.TokenConfigure](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		//cut by 10^decimals
		tokenUpdate = int64(params.TokenAmount * math.Pow(10, float64(tokenConfigure.Decimals)))
	}

	changeset := map[string]int64{
		WALLETS_KEY_GOLD:   goldUpdate,
		WALLETS_KEY_TOKENS: tokenUpdate,
	}

	_, _, err := nk.WalletUpdate(ctx, params.UserId, changeset, params.Metadata, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
