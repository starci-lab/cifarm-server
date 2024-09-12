package rpcs_nfts

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("update_premium_tile_nfts", UpdatePremiumTileNftsRpc)
	if err != nil {
		return err
	}
	return nil
}
