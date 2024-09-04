package entities

import (
	_constants "cifarm-server/src/constants"
	_farming_tools "cifarm-server/src/storage/farming_tools"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupFarmingTools(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	farming_tools := []_collections.FarmingTool{
		{
			Id: _constants.FARMING_TOOL_FERTILIZER_1,
		},
		{
			Id: _constants.FARMING_TOOL_FERTILIZER_2,
		},
		{
			Id: _constants.FARMING_TOOL_PESTICIDE,
		},
	}

	err := _farming_tools.WriteFarmingToolObjects(ctx, logger, db, nk, farming_tools)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
