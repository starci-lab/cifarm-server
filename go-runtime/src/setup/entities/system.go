package setup_entities

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupSystemUsers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	//write users
	object, err := collections_system.ReadUsers(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object != nil {
		return nil
	}

	users := collections_system.Users{
		UserIds: nil,
	}

	err = collections_system.WriteUsers(ctx, logger, db, nk, collections_system.WriteUsersParams{
		Users: users,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func SetupSystemActivityExperiences(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	err := collections_system.WriteActivityExperiences(ctx, logger, db, nk, collections_system.WriteActivityExperiencesParams{
		ActivityExperiences: collections_system.ActivityExperiences{
			Water:            3,
			UsePestiside:     3,
			UseFertilizer:    3,
			UseHerbicide:     3,
			HelpUseHerbicide: 3,
			HelpUsePestiside: 3,
			HelpWater:        3,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
