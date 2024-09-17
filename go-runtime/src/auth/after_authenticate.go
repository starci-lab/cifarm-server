package auth

import (
	collections_buildings "cifarm-server/src/collections/buildings"
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

	chain := in.Account.Vars["chainKey"]
	address := in.Account.Vars["accountAddress"]
	network := in.Account.Vars["network"]

	object, err := collections_config.ReadMetadata(ctx, logger, db, nk, collections_config.ReadMetadataParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		//first time login
		err = collections_config.WriteMetadata(ctx, logger, db, nk,
			collections_config.WriteMetadataParams{
				Metadata: collections_config.Metadata{
					ChainKey:       chain,
					AccountAddress: address,
					Network:        network,
				},
				UserId: userId,
			})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = collections_config.WritePlayerStats(ctx, logger, db, nk, collections_config.WritePlayerStatsParams{
			UserId: userId,
			PlayerStats: collections_config.PlayerStats{
				Level:       1,
				Experiences: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		positions := []collections_placed_items.Position{
			{X: 0, Y: -1},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: -1},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
		}

		var placedItems []collections_placed_items.PlacedItem
		for _, position := range positions {
			placedItems = append(placedItems, collections_placed_items.PlacedItem{
				ReferenceKey: collections_tiles.KEY_STARTER,
				Position:     position,
				Type:         collections_placed_items.TYPE_TILE,
				IsPlanted:    false,
			})
		}
		placedItems = append(placedItems, collections_placed_items.PlacedItem{
			ReferenceKey: collections_buildings.KEY_HOME,
			Position: collections_placed_items.Position{
				X: 0,
				Y: 5,
			},
			Type: collections_placed_items.TYPE_BUILDING,
		})

		err = collections_placed_items.WriteMany(ctx, logger, db, nk, collections_placed_items.WriteManyParams{
			PlacedItems: placedItems,
			UserId:      userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
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

		object, err := collections_system.ReadUsers(ctx, logger, db, nk)
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
		err = collections_system.WriteUsers(ctx, logger, db, nk, collections_system.WriteUsersParams{
			Users: *users,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	return nil
}
