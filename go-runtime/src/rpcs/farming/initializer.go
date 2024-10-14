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

	err = initializer.RegisterRpc("harvest_crop", HarvestCropRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("collect_animal_product", CollectAnimalProductRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("water", WaterRpc)
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

	err = initializer.RegisterRpc("feed_animal", FeedAnimalRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("use_fertilizer", UseFertilizerRpc)
	if err != nil {
		return err
	}
	return nil
}
