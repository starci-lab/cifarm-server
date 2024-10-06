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

type WriteMatchInfoParams struct {
	MatchInfo MatchInfo `json:"matchInfo"`
}

func WriteMatchInfo(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteMatchInfoParams,
) error {
	value, err := json.Marshal(params.MatchInfo)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_MATCH_INFO,
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

type WriteSpeedUpParams struct {
	SpeedUp SpeedUp `json:"speedUp"`
}

func WriteSpeedUp(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteSpeedUpParams,
) error {
	value, err := json.Marshal(params.SpeedUp)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_SPEEDUP,
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

type WriteActivityExperiencesParams struct {
	ActivityExperiences ActivityExperiences `json:"activityExperiences"`
}

func WriteActivityExperiences(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteActivityExperiencesParams,
) error {
	value, err := json.Marshal(params.ActivityExperiences)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_ACTIVITY_EXPERIENCES,
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

type WriteRewardsParams struct {
	Rewards Rewards `json:"rewards"`
}

func WriteRewards(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteRewardsParams,
) error {
	value, err := json.Marshal(params.Rewards)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_REWARDS,
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
