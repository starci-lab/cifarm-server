package matches_central

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type Match struct{}

type MatchState struct {
	Presences map[string]runtime.Presence `json:"-"`
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
		delete(matchState.Presences, presences[i].GetSessionId())
	}

	return matchState
}

type WrappedPlacedItems struct {
	PlacedItems []*collections_placed_items.PlacedItem `json:"placedItems"`
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
			objects, err := collections_placed_items.ReadMany(ctx, logger, db, nk, collections_placed_items.ReadsParams{
				UserId: presence.GetUserId(),
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
			err = dispatcher.BroadcastMessage(OP_CODE_PLACED_ITEMS_STATE, data, []runtime.Presence{
				presence,
			}, nil, true)
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
	return state, ""
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return state
}
