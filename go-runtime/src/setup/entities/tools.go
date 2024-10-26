package setup_entities

import (
	collections_tools "cifarm-server/src/collections/tools"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupTools(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	tools := []collections_tools.Tool{
		{
			Key:         collections_tools.KEY_SCYTHE,
			AvailableIn: collections_tools.AVAILABLE_IN_HOME,
		},
		{
			Key:         collections_tools.KEY_STEAL,
			AvailableIn: collections_tools.AVAILABLE_IN_NEIGHBOR,
		},
		{
			Key:         collections_tools.KEY_WATERCAN,
			AvailableIn: collections_tools.AVAILABLE_IN_BOTH,
		},
		{
			Key:         collections_tools.KEY_HERBICIDE,
			AvailableIn: collections_tools.AVAILABLE_IN_BOTH,
		},
		{
			Key:         collections_tools.KEY_PESTICIDE,
			AvailableIn: collections_tools.AVAILABLE_IN_BOTH,
		},
	}

	err := collections_tools.WriteMany(ctx, logger, db, nk, collections_tools.WriteManyParams{
		Tools: tools,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
