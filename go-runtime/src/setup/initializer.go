package setup

import (
	setup_entities "cifarm-server/src/setup/entities"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule,
) error {
	err := setup_entities.Initialize(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
