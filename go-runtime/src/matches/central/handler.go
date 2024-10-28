package matches_central

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type Match struct{}

type MatchState struct {
	Presences map[string]runtime.Presence `json:"-,omitempty"`
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		Presences: map[string]runtime.Presence{},
	}
	tickRate := 1
	label := LABEL
	return state, tickRate, label
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	lobbyState, ok := state.(*MatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	for i := 0; i < len(presences); i++ {
		lobbyState.Presences[presences[i].GetSessionId()] = presences[i]
	}

	return lobbyState
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	return state, true, ""
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	matchState, ok := state.(*MatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	for i := 0; i < len(presences); i++ {
		//when user leave, reset their visit state
		err := collections_player.WriteVisitState(ctx, logger, db, nk, collections_player.WriteVisitStateParams{
			VisitState: collections_player.VisitState{
				UserId: "",
			},
			UserId: presences[i].GetUserId(),
		})

		if err != nil {
			logger.Error(err.Error())
			return err
		}
		delete(matchState.Presences, presences[i].GetSessionId())
	}

	return matchState
}

type WrappedPlacedItems struct {
	PlacedItems []*collections_placed_items.PlacedItem `json:"placedItems,omitempty"`
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	matchState, ok := state.(*MatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	for _, presence := range matchState.Presences {
		go func() error {
			err := BroadcastPlacedItems(ctx, logger, db, nk, BroadcastPlacedItemsParams{
				presence:   presence,
				dispatcher: dispatcher,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			return nil
		}()
		go func() error {
			err := BroadcastPlayerStats(ctx, logger, db, nk, BroadcastPlayerStatsParams{
				presence:   presence,
				dispatcher: dispatcher,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			return nil
		}()
	}

	return state
}

func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	matchState, ok := state.(*MatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return state, ""
	}

	var presence runtime.Presence
	for _, _presence := range matchState.Presences {
		if _presence.GetUserId() == data {
			presence = _presence
			break
		}
	}

	err := BroadcastPlacedItems(ctx, logger, db, nk, BroadcastPlacedItemsParams{
		presence:   presence,
		dispatcher: dispatcher,
	})
	if err != nil {
		logger.Error(err.Error())
		return err, ""
	}
	return state, ""
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return state
}

type BroadcastPlacedItemsParams struct {
	presence   runtime.Presence
	dispatcher runtime.MatchDispatcher
}

func BroadcastPlacedItems(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params BroadcastPlacedItemsParams) error {
	object, err := collections_player.ReadVisitState(ctx, logger, db, nk, collections_player.ReadVisitStateParams{
		UserId: params.presence.GetUserId(),
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var currentUserId string
	if object != nil {
		visitState, err := collections_common.ToValue[collections_player.VisitState](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		currentUserId = visitState.UserId
	}

	if currentUserId == "" {
		currentUserId = params.presence.GetUserId()
	}

	objects, err := collections_placed_items.ReadMany(ctx, logger, db, nk, collections_placed_items.ReadManyParams{
		UserId: currentUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	values, err := collections_common.ToValues2[collections_placed_items.PlacedItem](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	wrapped := WrappedPlacedItems{
		PlacedItems: values,
	}

	data, err := json.Marshal(wrapped)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = params.dispatcher.BroadcastMessage(OP_CODE_PLACED_ITEMS_STATE, data, []runtime.Presence{
		params.presence,
	}, nil, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type BroadcastPlayerStatsParams struct {
	presence   runtime.Presence
	dispatcher runtime.MatchDispatcher
}

func BroadcastPlayerStats(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params BroadcastPlayerStatsParams) error {
	object, err := collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: params.presence.GetUserId(),
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	playerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	data, err := json.Marshal(playerStats)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = params.dispatcher.BroadcastMessage(OP_CODE_PLAYER_STATS, data, []runtime.Presence{
		params.presence,
	}, nil, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
