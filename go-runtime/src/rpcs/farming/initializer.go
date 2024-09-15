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

	err = initializer.RegisterRpc("harvest", HarvestRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("use_herbicide", UseHerbicideRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("use_pestiside", UsePestisideRpc)
	if err != nil {
		return err
	}
	return nil
}
