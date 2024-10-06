package matches_central

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

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
		//when user leave, reset their visit state
		err := collections_config.WriteVisitState(ctx, logger, db, nk, collections_config.WriteVisitStateParams{
			VisitState: collections_config.VisitState{
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
	}

	//broadcast next deli time
	var presences []runtime.Presence
	for _, presence := range matchState.Presences {
		presences = append(presences, presence)
	}
	err := BroadcastNextDeliveryTime(ctx, logger, db, nk, BroadcastNextDeliveryTimeParams{
		presences:  presences,
		dispatcher: dispatcher,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
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
	object, err := collections_config.ReadVisitState(ctx, logger, db, nk, collections_config.ReadVisitStateParams{
		UserId: params.presence.GetUserId(),
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var currentUserId string
	if object != nil {
		visitState, err := collections_common.ToValue[collections_config.VisitState](ctx, logger, db, nk, object)
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

type NextDeliveryTime struct {
	Time int64 `json:"time"`
}

type BroadcastNextDeliveryTimeParams struct {
	presences  []runtime.Presence
	dispatcher runtime.MatchDispatcher
}

func BroadcastNextDeliveryTime(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params BroadcastNextDeliveryTimeParams) error {
	now := time.Now()
	nextInterval := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())

	nextDeliveryTime := NextDeliveryTime{
		Time: nextInterval.Unix() - now.Unix(),
	}

	data, err := json.Marshal(nextDeliveryTime)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = params.dispatcher.BroadcastMessage(OP_CODE_NEXT_DELIVERY_TIME, data, params.presences, nil, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
