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
	Activities Activities `json:"activities"`
}

func WriteActivities(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteActivityExperiencesParams,
) error {
	value, err := json.Marshal(params.Activities)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_ACTIVITIES,
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

type WriteCropRandomnessParams struct {
	CropRandomness CropRandomness `json:"cropRandomness"`
}

func WriteCropRandomness(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteCropRandomnessParams,
) error {
	value, err := json.Marshal(params.CropRandomness)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_CROP_RANDOMNESS,
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

type WriteStarterConfigureParams struct {
	StarterConfigure StarterConfigure `json:"starterConfigure"`
}

func WriteStarterConfigure(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteStarterConfigureParams,
) error {
	value, err := json.Marshal(params.StarterConfigure)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_STARTER_CONFIGURE,
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

type WriteTokenConfigureParams struct {
	TokenConfigure TokenConfigure `json:"tokenConfigure"`
}

func WriteTokenConfigure(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteTokenConfigureParams,
) error {
	value, err := json.Marshal(params.TokenConfigure)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_TOKEN_CONFIGURE,
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

type WriteSpinConfigureParams struct {
	SpinConfigure SpinConfigure `json:"spinConfigure"`
}

func WriteSpinConfigure(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteSpinConfigureParams,
) error {
	value, err := json.Marshal(params.SpinConfigure)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_SPIN_CONFIGURE,
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
