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
	err := collections_system.WriteActivities(ctx, logger, db, nk, collections_system.WriteActivityExperiencesParams{
		Activities: collections_system.Activities{
			Water: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			UsePestiside: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			FeedAnimal: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			UseFertilizer: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			UseHerbicide: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			HelpUseHerbicide: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			HelpUsePestiside: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			HelpWater: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			ThiefCrop: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			HelpFeedAnimal: collections_system.ActivityInfo{
				ExperiencesGain: 50,
				EnergyCost:      1,
			},
			ThiefAnimalProduct: collections_system.ActivityInfo{
				ExperiencesGain: 3,
				EnergyCost:      1,
			},
			HelpUseFertilizer: collections_system.ActivityInfo{
				ExperiencesGain: 50,
				EnergyCost:      1,
			},
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func SetupSystemRewards(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	err := collections_system.WriteRewards(ctx, logger, db, nk, collections_system.WriteRewardsParams{
		Rewards: collections_system.Rewards{
			FromInvites: collections_system.FromInvites{
				Metrics: map[int]collections_system.Metric{
					1: {
						Key:   1,
						Value: 500,
					},
					2: {
						Key:   3,
						Value: 1000,
					},
					3: {
						Key:   10,
						Value: 2000,
					},
					4: {
						Key:   25,
						Value: 5000,
					},
				},
			},
			Referred: 200,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func SetupGlobalConstants(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	globalConstants := collections_system.GlobalConstants{
		GameRandomness: collections_system.GameRandomness{
			Theif3:            0.95,
			Theif2:            0.8,
			NeedWater:         0.5,
			IsWeedyOrInfested: 1,
		},
		Decimals:  5,
		GoldStart: 500,
	}

	err := collections_system.WriteGlobalConstants(ctx, logger, db, nk, collections_system.WriteGlobalConstantsParams{
		GlobalConstants: globalConstants,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
