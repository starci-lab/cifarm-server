package matches_central

import (
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterMatch(NAME,
		func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
			return &Match{}, nil
		})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	matchId, err := nk.MatchCreate(ctx, NAME, map[string]interface{}{})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = collections_system.WriteCentralMatchInfo(ctx, logger, db, nk, collections_system.WriteCentralMatchInfoParams{
		CentralMatchInfo: collections_system.CentralMatchInfo{
			MatchId: matchId,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(`Match started with id: %s`, matchId)
	return nil
}
