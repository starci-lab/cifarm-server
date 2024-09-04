package entities

import (
	_constants "cifarm-server/src/constants"
	_farming_tiles "cifarm-server/src/storage/farming_tiles"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupFarmingTiles(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	farming_tiles := []_collections.FarmingTile{
		{
			Id:           _constants.FARMING_TILE_BASIC_FARMING_TILE_STARTER,
			Price:        0,
			MaxOwnership: 6,
		},
		{
			Id:           _constants.FARMING_TILE_BASIC_FARMING_TILE_1,
			Price:        1000,
			MaxOwnership: 10,
		},
		{
			Id:           _constants.FARMING_TILE_BASIC_FARMING_TILE_2,
			Price:        2500,
			MaxOwnership: 30,
		},
		{
			Id:           _constants.FARMING_TILE_BASIC_FARMING_TILE_3,
			Price:        10000,
			MaxOwnership: 99999,
		},
	}

	err := _farming_tiles.WriteFarmingTilesObjects(ctx, logger, db, nk, _farming_tiles.WriteFarmingTileObjectsParams{
		FarmingTiles: farming_tiles,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
