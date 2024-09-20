package matches

import (
	collections_system "cifarm-server/src/collections/system"
	matches_central "cifarm-server/src/matches/central"
	matches_timer "cifarm-server/src/matches/timer"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := matches_central.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = matches_timer.Initialize(ctx, logger, db, nk, initializer)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	centralMatchId, err := nk.MatchCreate(ctx, matches_central.NAME, map[string]interface{}{})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(`Central match started with id: %s`, centralMatchId)

	timerMatchId, err := nk.MatchCreate(ctx, matches_timer.NAME, map[string]interface{}{})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(`Timer match started with id: %s`, timerMatchId)

	err = collections_system.WriteMatchInfo(ctx, logger, db, nk, collections_system.WriteMatchInfoParams{
		MatchInfo: collections_system.MatchInfo{
			CentralMatchId: centralMatchId,
			TimerMatchId:   timerMatchId,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
