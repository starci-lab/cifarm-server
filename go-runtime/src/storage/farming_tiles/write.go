package farming_tiles

import (
	"cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteFarmingTileObjectsParams struct {
	FarmingTiles []_collections.FarmingTile
}

func WriteFarmingTilesObjects(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteFarmingTileObjectsParams,
) error {
	var writes []*runtime.StorageWrite
	for _, farmingTile := range params.FarmingTiles {
		value, err := json.Marshal(farmingTile)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      constants.COLLECTION_FARMING_TILES,
			Key:             farmingTile.Id,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
