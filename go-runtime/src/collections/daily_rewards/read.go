package collections_daily_rewards

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadByKeyParams struct {
	Key string `json:"key"`
}

func ReadByKey(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadByKeyParams,
) (*api.StorageObject, error) {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{
		{
			Collection: COLLECTION_NAME,
			Key:        params.Key,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects) == 0 {
		return nil, nil
	}

	object := objects[0]
	return object, nil
}

type ReadHighestPossibleDayParams struct {
	UserId string `json:"userId"`
	Streak int    `json:"streak"`
}

func ReadHighestPossibleDay(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadHighestPossibleDayParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_HIGHEST_POSSIBLE_DAY
	query := fmt.Sprintf("+user_id:%s +value.day:<=%v",
		params.UserId, params.Streak)
	order := []string{
		"-value.day",
	}

	objects, err := nk.StorageIndexList(ctx, "", name, query, 1, order)
	logger.Debug("%v", len(objects.Objects))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if len(objects.Objects) == 0 {
		return nil, nil
	}
	object := objects.Objects[0]
	return object, nil
}
