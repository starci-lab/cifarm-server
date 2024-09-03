package daily_reward

import (
	_constants "cifarm-server/src/constants"
	_wallets "cifarm-server/src/utils/wallets"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteDailyRewardObjectParams struct {
	Amount int64 `json:"amount"`
	Days   int   `json:"days"`
}

func WriteDailyRewardObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteDailyRewardObjectParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	value, err := json.Marshal(params)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      _constants.COLLECTION_REWARDS,
			Key:             uuid.NewString(),
			UserID:          userId,
			Value:           string(value),
			PermissionRead:  1,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func ReadLatestDailyRewardObject(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) (*api.StorageObject, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	name := _constants.STORAGE_INDEX_LATEST_DAILY_REWARD_OBJECTS
	query := fmt.Sprintf("user_id:%s", userId)
	order := []string{
		"-create_time",
	}

	dailyRewards, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(dailyRewards.Objects) == 0 {
		return nil, nil
	}
	var latest = dailyRewards.Objects[0]
	return latest, nil
}

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

type DailyRewardObjectValue struct {
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
	object, err := ReadLatestDailyRewardObject(ctx, logger, db, nk)
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
		err = WriteDailyRewardObject(ctx, logger, db, nk, WriteDailyRewardObjectParams{
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
	var value *DailyRewardObjectValue
	err = json.Unmarshal([]byte(object.Value), &value)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	days := value.Days
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

	err = WriteDailyRewardObject(ctx, logger, db, nk, WriteDailyRewardObjectParams{
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
