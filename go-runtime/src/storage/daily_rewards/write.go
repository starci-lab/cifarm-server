package daily_rewards

import (
	_constants "cifarm-server/src/constants"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
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
