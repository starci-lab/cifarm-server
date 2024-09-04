package farming_tiles

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageIndexFarmingTiles(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := _constants.STORAGE_INDEX_FARMING_TILES
	collection := _constants.COLLECTION_FARMING_TILES
	key := ""
	fields := []string{
		"price",
	}
	sortableFields := []string{}
	maxEntries := 100
	indexOnly := false
	err := initializer.RegisterStorageIndex(name, collection, key, fields, sortableFields, maxEntries, indexOnly)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
