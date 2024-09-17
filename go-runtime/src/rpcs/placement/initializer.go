package rpcs_placement

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("move", MoveRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("place_tile", PlaceTileRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("recover_tile", RecoverTileRpc)
	if err != nil {
		return err
	}

	return nil
}
