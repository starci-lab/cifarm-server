package collections_player

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteMetadataParams struct {
	Metadata Metadata `json:"metadata"`
	UserId   string   `json:"userId"`
}

func WriteMetadata(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteMetadataParams,
) error {
	value, err := json.Marshal(params.Metadata)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_METADATA,
			UserID:          params.UserId,
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

type WriteVisitStateParams struct {
	VisitState VisitState `json:"visitState"`
	UserId     string     `json:"userId"`
}

func WriteVisitState(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteVisitStateParams,
) error {
	value, err := json.Marshal(params.VisitState)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_VISIT_STATE,
			UserID:          params.UserId,
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

type WritePlayerStatsParams struct {
	PlayerStats PlayerStats `json:"playerStats"`
	UserId      string      `json:"userId"`
}

func WritePlayerStats(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WritePlayerStatsParams,
) error {
	value, err := json.Marshal(params.PlayerStats)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_PLAYER_STATS,
			UserID:          params.UserId,
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

type WriteRewardTrackerParams struct {
	RewardTracker RewardTracker `json:"rewardTracker"`
	UserId        string        `json:"userId"`
}

func WriteRewardTracker(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteRewardTrackerParams,
) error {
	value, err := json.Marshal(params.RewardTracker)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_REWARD_TRACKER,
			UserID:          params.UserId,
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

type WriteFollowingsParams struct {
	Followings Followings `json:"followings"`
	UserId     string     `json:"userId"`
}

func WriteFollowings(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteFollowingsParams,
) error {
	value, err := json.Marshal(params.Followings)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             KEY_FOLLOWINGS,
			UserID:          params.UserId,
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

type IncreaseExperiencesParams struct {
	Amount int    `json:"amount"`
	UserId string `json:"userId"`
}

func IncreaseExperiences(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params IncreaseExperiencesParams) error {
	object, err := ReadPlayerStats(ctx, logger, db, nk, ReadPlayerStatsParams{
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	playerStats, err := collections_common.ToValue[PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	//each level need 50 exp, next level need 20 more
	//50=>100=>150=>...
	totalExperiences := playerStats.LevelInfo.Experiences + params.Amount
	for {
		if totalExperiences >= playerStats.LevelInfo.ExperienceQuota {
			totalExperiences -= playerStats.LevelInfo.ExperienceQuota
			playerStats.LevelInfo.Level += 1
			playerStats.LevelInfo.ExperienceQuota = playerStats.LevelInfo.Level * 50
			//max energy increase by quota, current energy restore to max
			playerStats.EnergyInfo.MaxEnergy += playerStats.EnergyInfo.EnergyQuota
			playerStats.EnergyInfo.CurrentEnergy = playerStats.EnergyInfo.MaxEnergy
		} else {
			playerStats.LevelInfo.Experiences = totalExperiences
			break
		}
	}

	err = WritePlayerStats(ctx, logger, db, nk, WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type DecreaseEnergyParams struct {
	Amount int    `json:"amount"`
	UserId string `json:"userId"`
}

func DecreaseEnergy(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DecreaseEnergyParams) error {
	object, err := ReadPlayerStats(ctx, logger, db, nk, ReadPlayerStatsParams{
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	playerStats, err := collections_common.ToValue[PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if (playerStats.EnergyInfo.CurrentEnergy - params.Amount) < 0 {
		errMsg := "not enough energy"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	playerStats.EnergyInfo.CurrentEnergy -= params.Amount
	err = WritePlayerStats(ctx, logger, db, nk, WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
