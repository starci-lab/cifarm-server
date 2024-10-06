package rpcs_daily_rewards

import (
	collections_common "cifarm-server/src/collections/common"
	collections_daily_rewards "cifarm-server/src/collections/daily_rewards"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func CanUserClaimDailyReward(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (bool, error) {
	if object == nil {
		errMsg := "object is nil"
		return false, errors.New(errMsg)
	}

	objectCreateTime := time.Unix(object.CreateTime.Seconds, 0).UTC()
	startOfToday := time.Date(
		objectCreateTime.Year(),
		objectCreateTime.Month(),
		objectCreateTime.Day(),
		0,
		0,
		0,
		0,
		time.UTC)
	//startOfTomorrow := startOfToday
	startOfTomorrow := startOfToday.Add(24 * time.Hour)
	now := time.Now().UTC().Unix()

	result := now >= startOfTomorrow.Unix()
	return result, nil
}

type ClaimDailyRewardRpcResponse struct {
	Amount int64 `json:"amount"`
	Days   int   `json:"days"`
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

	object, err := collections_daily_rewards.ReadLatest(ctx, logger, db, nk, collections_daily_rewards.ReadLatestParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		amount := int64(100)
		days := 1
		err := wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
			Amount: amount,
			UserId: userId,
			Metadata: map[string]interface{}{
				"name": "Daily reward",
				"days": days,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		err = collections_daily_rewards.Write(ctx, logger, db, nk, collections_daily_rewards.WriteParams{
			UserId: userId,
			DailyReward: collections_daily_rewards.DailyReward{
				Amount: amount,
				Days:   days,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		value, err := json.Marshal(ClaimDailyRewardRpcResponse{
			Amount: amount,
			Days:   days,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		return string(value), err
	}

	can, err := CanUserClaimDailyReward(ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if !can {
		errMsg := "you have claimed reward today"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	amount := int64(100)
	dailyReward, err := collections_common.ToValue[collections_daily_rewards.DailyReward](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	days := dailyReward.Days
	days++

	err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
		Amount: amount,
		UserId: userId,
		Metadata: map[string]interface{}{
			"name": "Daily reward",
			"days": days,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = collections_daily_rewards.Write(ctx, logger, db, nk, collections_daily_rewards.WriteParams{
		DailyReward: collections_daily_rewards.DailyReward{
			Amount: amount,
			Days:   days,
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	_value, err := json.Marshal(ClaimDailyRewardRpcResponse{
		Amount: amount,
		Days:   days,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
