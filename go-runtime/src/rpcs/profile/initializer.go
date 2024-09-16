package rpcs_profiles

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("list_inventories", ListInventoriesRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("deliver_products", DeliverProductsRpc)
	if err != nil {
		return err
	}
	return nil
}
