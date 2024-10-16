package rpcs_community

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("visit", VisitRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("thief_crop", ThiefCropRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("search_users", SearchUserRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("get_random_user", GetRandomUserRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("return", ReturnRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("help_use_herbicide", HelpUseHerbicideRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("help_use_pestiside", HelpUsePestisideRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("help_water", HelpWaterRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("thief_animal_product", ThiefAnimalProductRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("follow", FollowRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("unfollow", UnfollowRpc)
	if err != nil {
		return err
	}

	return nil
}
