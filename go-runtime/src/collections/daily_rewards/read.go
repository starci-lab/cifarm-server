package collections_daily_rewards

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type ReadLatestParams struct {
	UserId string `json:"userId"`
}

func ReadLatest(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params ReadLatestParams,
) (*api.StorageObject, error) {
	name := STORAGE_INDEX_LATEST
	query := ""
	order := []string{
		"-create_time",
	}

	objects, err := nk.StorageIndexList(ctx, params.UserId, name, query, 1, order)
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
