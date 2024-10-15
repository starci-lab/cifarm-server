package rpcs_daily_rewards

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand/v2"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type CanUserClaimDailyRewardParams struct {
	UserId string `json:"userId"`
}

type ClaimDailyRewardRpcResponse struct {
	//response only neccessary for the lasted date
	LastDailyRewardPossibility collections_daily_rewards.LastDailyRewardPossibility `json:"lastDailyRewardPossibility"`
}

func ClaimDailyRewardRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	var lastDailyRewardPossibility collections_daily_rewards.LastDailyRewardPossibility

	//get last claimed
	object, err := collections_config.ReadPlayerStats(ctx, logger, db, nk, collections_config.ReadPlayerStatsParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//check claim possible
	playerStats, err := collections_common.ToValue[collections_config.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	lastClaimedDateBegin := time.Unix(playerStats.DailyRewardsInfo.LastClaimedTime, 0).UTC()
	startOfLastClaimedDate := time.Date(
		lastClaimedDateBegin.Year(),
		lastClaimedDateBegin.Month(),
		lastClaimedDateBegin.Day(),
		0,
		0,
		0,
		0,
		time.UTC)

	tomorrowAfterLastClaimedDate := startOfLastClaimedDate.Add(24 * time.Hour)
	now := time.Now().UTC().Unix()

	result := now >= tomorrowAfterLastClaimedDate.Unix()
	if !result {
		errMsg := "you have already claimed the daily reward this day"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process logic
	//if you do not claim the reward for 2 days, the streak will be reset
	if now > tomorrowAfterLastClaimedDate.Add(24*time.Hour).Unix() {
		playerStats.DailyRewardsInfo.Streak = 0
	} else {
		playerStats.DailyRewardsInfo.Streak++
	}

	//update the last claimed time
	playerStats.DailyRewardsInfo.LastClaimedTime = now
	playerStats.DailyRewardsInfo.NumberOfClaims++

	//write the player stats
	err = collections_config.WritePlayerStats(ctx, logger, db, nk, collections_config.WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check best reward
	object, err = collections_daily_rewards.ReadHighestPossibleDay(ctx, logger, db, nk, collections_daily_rewards.ReadHighestPossibleDayParams{
		UserId: userId,
		Streak: playerStats.DailyRewardsInfo.Streak,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "daily reward not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	dailyReward, err := collections_common.ToValue[collections_daily_rewards.DailyReward](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if !dailyReward.IsLastDay {
		//only add golds if it is not the last day
		err := wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     userId,
			GoldAmount: dailyReward.Amount,
			Metadata: map[string]interface{}{
				"name": "Daily reward",
				"time": time.Now().Format(time.RFC850),
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
	} else {
		//randomize the reward
		randomValue := rand.Float64()
		for _, dailyRewardPossibility := range dailyReward.DailyRewardPossibilities {
			if randomValue <= dailyRewardPossibility.ThresholdMax && randomValue > dailyRewardPossibility.ThresholdMin {
				lastDailyRewardPossibility = dailyRewardPossibility
				break
			}
		}
		err := wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:      userId,
			GoldAmount:  lastDailyRewardPossibility.GoldAmount,
			TokenAmount: lastDailyRewardPossibility.TokenAmount,
			Metadata: map[string]interface{}{
				"name": "Last daily reward",
				"time": time.Now().Format(time.RFC850),
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
	}

	_value, err := json.Marshal(ClaimDailyRewardRpcResponse{
		LastDailyRewardPossibility: lastDailyRewardPossibility,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
