package rpcs_shop

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("buy_seeds", BuySeedsRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("buy_supplies", BuySuppliesRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("buy_tile", BuyTileRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("construct_building", ConstructBuildingRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("buy_animal", BuyAnimalRpc)
	if err != nil {
		return err
	}
	return nil
}
