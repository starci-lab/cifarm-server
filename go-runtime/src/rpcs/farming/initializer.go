package rpcs_farming

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("plant_seed", PlantSeedRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("place_tile", PlaceTileRpc)
	if err != nil {
		return err
	}
	return nil
}
