package auth

import (
	_constants "cifarm-server/src/constants"
	_config "cifarm-server/src/storage/config"
	_inventories "cifarm-server/src/storage/inventories"
	_collections "cifarm-server/src/types/collections"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Claims struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func AfterAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	out *api.Session,
	in *api.AuthenticateCustomRequest,
) error {
	chain := in.Account.Vars["chain"]
	address := in.Account.Vars["address"]

	object, err := _config.ReadConfigPlayerMetdataObject(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	config, err := _config.ToConfigPlayerMetdata(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if config == nil {
		err = _config.WriteConfigPlayerMetdataObject(ctx, logger, db, nk, _config.WriteConfigPlayerMetdataObjectParams{
			PlayerMetadata: _collections.PlayerMetadata{
				Chain:   chain,
				Address: address,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		err = _inventories.WriteInventoryObject(ctx, logger, db, nk, _inventories.WriteInventoryObjectParams{
			Id:       _constants.FARMING_TILE_BASIC_FARMING_TILE_STARTER,
			Type:     _constants.TYPE_FARMING_TILE,
			Quantity: 6,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
			Amount: 500,
			Metadata: map[string]interface{}{
				"name": "Initial",
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	return nil
}
