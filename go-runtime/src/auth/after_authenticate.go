package auth

import (
	_constants "cifarm-server/src/constants"
	_config "cifarm-server/src/storage/config"
	_placed_items "cifarm-server/src/storage/placed_items"
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
		err = _config.WriteConfigPlayerMetdataObject(ctx, logger, db, nk,
			_collections.PlayerMetadata{
				Chain:   chain,
				Address: address,
			})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		positions := []_collections.Position{
			{X: -0.5, Y: 0.5},
			{X: 0.5, Y: 0.5},
			{X: 1.5, Y: -0.5},
			{X: -0.5, Y: -0.5},
			{X: 0.5, Y: -0.5},
			{X: 1.5, Y: -0.5},
		}

		var placedItems []_collections.PlacedItem
		for _, pos := range positions {
			placedItems = append(placedItems, _collections.PlacedItem{
				Id:       _constants.FARMING_TILE_BASIC_FARMING_TILE_STARTER,
				Position: pos,
			})
		}

		err = _placed_items.WritePlacedItemObjects(ctx, logger, db, nk, placedItems)
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
