package rpcs_upgrades

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("upgrade_building", UpgradeBuildingRpc)
	if err != nil {
		return err
	}

	return nil
}
