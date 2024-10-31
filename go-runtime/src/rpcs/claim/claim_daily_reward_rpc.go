package rpcs_claim

import (
	collections_common "cifarm-server/src/collections/common"
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	collections_player "cifarm-server/src/collections/player"
	"cifarm-server/src/utils"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"math/rand/v2"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type ClaimDailyRewardRpcResponse struct {
	//response only neccessary for the lasted date
	LastDailyRewardPossibilityKey string `json:"lastDailyRewardPossibilityKey"`
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

	//get reward tracker
	object, err := collections_player.ReadRewardTracker(ctx, logger, db, nk, collections_player.ReadRewardTrackerParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		errMsg := "reward tracker not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//check claim possible
	rewardTracker, err := collections_common.ToValue[collections_player.RewardTracker](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	lastClaimDateBegin := time.Unix(rewardTracker.DailyRewardsInfo.LastClaimTime, 0).UTC()
	startOfLastClaimDate := utils.StartOfTomorow(lastClaimDateBegin)

	tomorrowAfterLastClaimDate := startOfLastClaimDate.Add(24 * time.Hour)
	now := time.Now().UTC().Unix()

	result := now >= tomorrowAfterLastClaimDate.Unix()
	if !result {
		errMsg := "you have already claimed the daily reward this day"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	//process logic
	//if you do not claim the reward for 2 days, the streak will be reset
	if now > tomorrowAfterLastClaimDate.Add(24*time.Hour).Unix() {
		rewardTracker.DailyRewardsInfo.Streak = 0
	} else {
		rewardTracker.DailyRewardsInfo.Streak++
	}

	//update the last claimed time
	rewardTracker.DailyRewardsInfo.LastClaimTime = now
	rewardTracker.DailyRewardsInfo.NumberOfClaims++

	//write the player stats
	err = collections_player.WriteRewardTracker(ctx, logger, db, nk, collections_player.WriteRewardTrackerParams{
		RewardTracker: *rewardTracker,
		UserId:        userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//check best reward
	object, err = collections_daily_rewards.ReadHighestPossibleDay(ctx, logger, db, nk, collections_daily_rewards.ReadHighestPossibleDayParams{
		MaxPossibleDay: rewardTracker.DailyRewardsInfo.Streak + 1,
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

	var lastDailyRewardPossibility collections_daily_rewards.LastDailyRewardPossibility
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
			if randomValue < dailyRewardPossibility.ThresholdMax && randomValue >= dailyRewardPossibility.ThresholdMin {
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
		LastDailyRewardPossibilityKey: lastDailyRewardPossibility.Key,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
