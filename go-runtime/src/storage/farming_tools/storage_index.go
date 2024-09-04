package farming_tools

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitializeStorageIndexFarmingTools(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	initializer runtime.Initializer,
) error {
	name := _constants.STORAGE_INDEX_FARMING_TOOLS
	collection := _constants.COLLECTION_FARMING_TOOLS
	key := ""
	fields := []string{
		"days",
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
