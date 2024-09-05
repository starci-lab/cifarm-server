package collections_system

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteUsersParams struct {
	Users Users `json:"users"`
}

func WriteUsers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteUsersParams,
) error {
	value, err := json.Marshal(params.Users)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_USERS,
			Value:           string(value),
			PermissionRead:  0,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func WriteLastServerUptime(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	lastServerUptime := LastServerUptime{
		TimeInSeconds: time.Now().Unix(),
	}
	value, err := json.Marshal(lastServerUptime)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_LAST_SERVER_UPTIME,
			Value:           string(value),
			PermissionRead:  0,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type WriteCentralMatchInfoParams struct {
	CentralMatchInfo CentralMatchInfo `json:"centralMatchInfo"`
}

func WriteCentralMatchInfo(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteCentralMatchInfoParams,
) error {
	value, err := json.Marshal(params.CentralMatchInfo)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_CENTRAL_MATCH_INFO,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
