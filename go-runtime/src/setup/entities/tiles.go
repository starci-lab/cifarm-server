package setup_entities

import (
	collections_tiles "cifarm-server/src/collections/tiles"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupTiles(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	tiles := []collections_tiles.Tile{
		{
			ReferenceId:  collections_tiles.KEY_STARTER,
			Price:        0,
			MaxOwnership: 6,
		},
		{
			ReferenceId:  collections_tiles.KEY_BASIC_1,
			Price:        1000,
			MaxOwnership: 10,
		},
		{
			ReferenceId:  collections_tiles.KEY_BASIC_2,
			Price:        2500,
			MaxOwnership: 30,
		},
		{
			ReferenceId:  collections_tiles.KEY_BASIC_3,
			Price:        10000,
			MaxOwnership: 9999,
		},
	}

	err := collections_tiles.WriteMany(ctx, logger, db, nk, collections_tiles.WriteManyParams{
		Tiles: tiles,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
