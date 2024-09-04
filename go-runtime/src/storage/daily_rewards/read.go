package daily_rewards

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

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
	name := _constants.STORAGE_INDEX_DAILY_REWARDS
	query := fmt.Sprintf(`+user_id:%s`, userId)
	order := []string{
		"-create_time",
	}

	objects, err := nk.StorageIndexList(ctx, userId, name, query, 1, order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}
	var latest = objects.Objects[0]
	return latest, nil
}

func ToLatestDailyReward(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	object *api.StorageObject,
) (*_collections.DailyReward, error) {
	var dailyReward *_collections.DailyReward
	if object == nil {
		return nil, nil
	}
	err := json.Unmarshal([]byte(object.Value), &dailyReward)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return dailyReward, nil
}
