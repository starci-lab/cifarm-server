package collections_daily_rewards

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	DailyReward DailyReward `json:"dailyReward"`
	UserId      string      `json:"userId"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) error {
	key := uuid.NewString()

	value, err := json.Marshal(params.DailyReward)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
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
