package rpcs_miscellaneous

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("fetch_central_instantly", FetchCentralInstantlyRpc)
	if err != nil {
		return err
	}
	return nil
}
