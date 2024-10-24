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
			Key:             collections_tiles.KEY_STARTER_TILE,
			Price:           0,
			MaxOwnership:    6,
			IsNFT:           false,
			AvailableInShop: true,
		},
		{
			Key:             collections_tiles.KEY_BASIC_TILE_1,
			Price:           1000,
			MaxOwnership:    10,
			IsNFT:           false,
			AvailableInShop: true,
		},
		{
			Key:             collections_tiles.KEY_BASIC_TILE_2,
			Price:           2500,
			MaxOwnership:    30,
			IsNFT:           false,
			AvailableInShop: true,
		},
		{
			Key:             collections_tiles.KEY_BASIC_TILE_3,
			Price:           10000,
			MaxOwnership:    9999,
			IsNFT:           false,
			AvailableInShop: true,
		},
		{
			Key:             collections_tiles.KEY_FERTILE_TILE,
			AvailableInShop: false,
			IsNFT:           true,
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
