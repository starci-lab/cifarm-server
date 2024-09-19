package collections_config

import (
	collections_common "cifarm-server/src/collections/common"
	"context"
	"database/sql"
	"encoding/json"

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

type IncreaseExperiencesParams struct {
	Amount int64  `json:"amount"`
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
	totalExperiences := playerStats.Experiences + int64(params.Amount)
	for {
		if totalExperiences >= playerStats.ExperienceQuota {
			totalExperiences -= playerStats.ExperienceQuota
			playerStats.Level += 1
			playerStats.ExperienceQuota = 50 + int64(playerStats.Level-1)*50
		} else {
			playerStats.Experiences = totalExperiences
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
