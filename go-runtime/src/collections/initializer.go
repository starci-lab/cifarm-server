package storage

import (
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_nfts "cifarm-server/src/collections/nfts"
	collections_placed_items "cifarm-server/src/collections/placed_items"
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

	err = collections_nfts.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
