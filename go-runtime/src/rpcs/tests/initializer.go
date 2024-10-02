package rpcs_tests

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("test_speed_up", SpeedUpRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("test_delivery", DeliveryRpc)
	if err != nil {
		return err
	}
	err = initializer.RegisterRpc("test_hack_gold", HackGoldRpc)
	if err != nil {
		return err
	}
	return nil
}
