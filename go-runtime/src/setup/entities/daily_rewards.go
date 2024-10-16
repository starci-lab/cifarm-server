package setup_entities

import (
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupDailyRewards(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	dailyRewards := []collections_daily_rewards.DailyReward{
		{
			Key:       collections_daily_rewards.KEY_DAY_1,
			Amount:    100,
			Day:       1,
			IsLastDay: false,
		},
		{
			Key:       collections_daily_rewards.KEY_DAY_2,
			Amount:    200,
			Day:       2,
			IsLastDay: false,
		},
		{
			Key:       collections_daily_rewards.KEY_DAY_3,
			Amount:    300,
			Day:       3,
			IsLastDay: false,
		},
		{
			Key:       collections_daily_rewards.KEY_DAY_4,
			Amount:    600,
			Day:       4,
			IsLastDay: false,
		},
		{
			Key:       collections_daily_rewards.KEY_DAY_5,
			Day:       5,
			IsLastDay: true,
			DailyRewardPossibilities: map[int]collections_daily_rewards.LastDailyRewardPossibility{
				1: {
					Key:          "possibility1",
					GoldAmount:   1000,
					ThresholdMin: 0,
					ThresholdMax: 0.8,
				},
				2: {
					Key:          "possibility2",
					GoldAmount:   1500,
					ThresholdMin: 0.8,
					ThresholdMax: 0.9,
				},
				3: {
					Key:          "possibility3",
					GoldAmount:   2000,
					ThresholdMin: 0.9,
					ThresholdMax: 0.95,
				},
				4: {
					Key:          "possibility4",
					TokenAmount:  3,
					ThresholdMin: 0.95,
					ThresholdMax: 0.99,
				},
				5: {
					Key:          "possibility5",
					TokenAmount:  10,
					ThresholdMin: 0.99,
					ThresholdMax: 1,
				},
			},
		},
	}

	err := collections_daily_rewards.WriteMany(ctx, logger, db, nk, collections_daily_rewards.WriteManyParams{
		DailyRewards: dailyRewards,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
