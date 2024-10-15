package collections

import (
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	collections_delivering_products "cifarm-server/src/collections/delivering_products"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	err := collections_inventories.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_placed_items.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_daily_rewards.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_player.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_delivering_products.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
