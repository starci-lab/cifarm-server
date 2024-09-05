package matches_central

import (
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type LobbyMatch struct{}

type LobbyMatchState struct {
	Presences map[string]runtime.Presence `json:"-"`
}

func (m *LobbyMatch) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &LobbyMatchState{
		Presences: map[string]runtime.Presence{},
	}
	tickRate := 1
	label := LABEL
	return state, tickRate, label
}

func (m *LobbyMatch) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	lobbyState, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil
	}

	for i := 0; i < len(presences); i++ {
		lobbyState.Presences[presences[i].GetSessionId()] = presences[i]
	}

	return lobbyState
}

func (m *LobbyMatch) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	_, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil, false, ""
	}
	return nil, false, ""
}

func (m *LobbyMatch) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	lobbyState, ok := state.(*LobbyMatchState)
	if !ok {
		logger.Error("state not a valid lobby state object")
		return nil
	}

	for i := 0; i < len(presences); i++ {
		delete(lobbyState.Presences, presences[i].GetSessionId())
	}

	return lobbyState
}

func (m *LobbyMatch) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	_, ok := state.(*LobbyMatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	return state
}

func (m *LobbyMatch) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return nil, ""
}

func (m *LobbyMatch) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return nil
}
