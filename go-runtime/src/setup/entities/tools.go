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
			ReferenceId: collections_tools.KEY_FERTILIZER_1,
		},
		{
			ReferenceId: collections_tools.KEY_FERTILIZER_2,
		},
		{
			ReferenceId: collections_tools.KEY_PESTICIDE,
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
