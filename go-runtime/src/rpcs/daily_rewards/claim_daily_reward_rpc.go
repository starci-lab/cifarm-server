package daily_rewards

import (
	_daily_rewards "cifarm-server/src/storage/daily_rewards"
	_wallets "cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type CanClaimDailyRewardRpcResponse struct {
	Amount int64 `json:"amount"`
	Days   int   `json:"days"`
}

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

func ClaimDailyRewardRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	object, err := _daily_rewards.ReadLatestDailyRewardObject(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if object == nil {
		amount := int64(500)
		days := 1
		err := _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
			Amount: amount,
			Metadata: map[string]interface{}{
				"name": "Daily reward",
				"days": days,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		err = _daily_rewards.WriteDailyRewardObject(ctx, logger, db, nk, _daily_rewards.WriteDailyRewardObjectParams{
			Amount: amount,
			Days:   days,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

		value, err := json.Marshal(CanClaimDailyRewardRpcResponse{
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
	dailyReward, err := _daily_rewards.ReadLatestDailyRewardObjectValue(
		ctx,
		logger,
		db,
		nk,
		object,
	)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	days := dailyReward.Days
	days++

	err = _wallets.UpdateWallet(ctx, logger, db, nk, _wallets.UpdateWalletParams{
		Amount: amount,
		Metadata: map[string]interface{}{
			"name": "Daily reward",
			"days": days,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	err = _daily_rewards.WriteDailyRewardObject(ctx, logger, db, nk, _daily_rewards.WriteDailyRewardObjectParams{
		Amount: amount,
		Days:   days,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	_value, err := json.Marshal(CanClaimDailyRewardRpcResponse{
		Amount: amount,
		Days:   days,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(_value), err
}
