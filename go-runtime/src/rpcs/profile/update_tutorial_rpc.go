package rpcs_profile

import (
	collections_common "cifarm-server/src/collections/common"
	collections_player "cifarm-server/src/collections/player"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdateTutorialRpcParams struct {
	TutorialIndex int `json:"tutorialIndex"`
	StepIndex     int `json:"stepIndex"`
}

func UpdateTutorialRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *UpdateTutorialRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err := collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}
	playerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	playerStats.TutorialInfo.TutorialIndex = params.TutorialIndex
	playerStats.TutorialInfo.StepIndex = params.StepIndex

	err = collections_player.WritePlayerStats(ctx, logger, db, nk, collections_player.WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return "", nil
}
