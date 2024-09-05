package auth

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_system "cifarm-server/src/collections/system"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"errors"

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
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	chain := in.Account.Vars["chain"]
	address := in.Account.Vars["address"]

	object, err := collections_config.ReadMetadataByKey(ctx, logger, db, nk, collections_config.ReadMetadataByKeyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	config, err := collections_common.ToValue[collections_config.Metadata](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if config == nil {
		err = collections_config.Write(ctx, logger, db, nk,
			collections_config.WriteParams{
				Metadata: collections_config.Metadata{
					Chain:   chain,
					Address: address,
				},
				UserId: userId,
			})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		positions := []collections_placed_items.Position{
			{X: -0.5, Y: 0.5},
			{X: 0.5, Y: 0.5},
			{X: 1.5, Y: -0.5},
			{X: -0.5, Y: -0.5},
			{X: 0.5, Y: -0.5},
			{X: 1.5, Y: -0.5},
		}

		var placedItems []collections_placed_items.PlacedItem
		for _, position := range positions {
			placedItems = append(placedItems, collections_placed_items.PlacedItem{
				ReferenceId: collections_tiles.KEY_STARTER,
				Position:    position,
				Type:        collections_placed_items.TYPE_TILE,
				IsPlanted:   false,
			})
		}

		err = collections_placed_items.WriteMany(ctx, logger, db, nk, collections_placed_items.WriteManyParams{
			PlacedItems: placedItems,
			UserId:      userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			Amount: 500,
			Metadata: map[string]interface{}{
				"name": "Initial",
			},
			UserId: userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		object, err := collections_system.ReadByKey(ctx, logger, db, nk)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		users, err := collections_common.ToValue[collections_system.Users](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		users.UserIds = append(users.UserIds, userId)
		err = collections_system.Write(ctx, logger, db, nk, collections_system.WriteParams{
			Users: *users,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	return nil
}
