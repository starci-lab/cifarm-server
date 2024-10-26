package matches_timer

import (
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/utils"
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
	PlacedItems []*collections_placed_items.PlacedItem `json:"placedItems"`
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	matchState, ok := state.(*MatchState)
	if !ok {
		errMsg := "state not a valid lobby state object"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	//broadcast next deli time
	var presences []runtime.Presence
	for _, presence := range matchState.Presences {
		presences = append(presences, presence)
		//for this loop, we only need to broadcast the user's cooldown timers
		err := BroadcaseUserCooldownTimers(ctx, logger, db, nk, BroadcaseUserCooldownTimersParams{
			presence:   presence,
			dispatcher: dispatcher,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	err := BroadcastGlobalCooldownTimers(ctx, logger, db, nk, BroadcastGlobalCooldownTimersParams{
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
	return state, ""
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	return state
}

type BroadcastPlacedItemsParams struct {
	Presence   runtime.Presence
	Dispatcher runtime.MatchDispatcher
}

type UserCooldownTimers struct {
	//spin
	NextFreeSpinCooldown int64 `json:"nextFreeSpinCooldown"`
	IsSpinFree           bool  `json:"isSpinFree"`

	//daily rewards
	NextDailyRewardCooldown int64 `json:"nextDailyRewardCooldown"`
	ClaimedThisDay          bool  `json:"claimedThisDay"`
}

type GlobalCooldownTimers struct {
	NextDeliveryCooldown int64 `json:"nextDeliveryCooldown"`
}

type BroadcastGlobalCooldownTimersParams struct {
	presences  []runtime.Presence
	dispatcher runtime.MatchDispatcher
}

func BroadcastGlobalCooldownTimers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params BroadcastGlobalCooldownTimersParams) error {
	now := time.Now()
	startOfNextDay := utils.StartOfTomorow(now)

	globalCooldownTimers := GlobalCooldownTimers{
		NextDeliveryCooldown: startOfNextDay.Unix() - now.Unix(),
	}

	data, err := json.Marshal(globalCooldownTimers)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = params.dispatcher.BroadcastMessage(OP_CODE_GLOBAL_COOLDOWN_TIMERS, data, params.presences, nil, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

type BroadcaseUserCooldownTimersParams struct {
	presence   runtime.Presence
	dispatcher runtime.MatchDispatcher
}

func BroadcaseUserCooldownTimers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params BroadcaseUserCooldownTimersParams) error {
	//reward tracker
	object, err := collections_player.ReadRewardTracker(ctx, logger, db, nk, collections_player.ReadRewardTrackerParams{
		UserId: params.presence.GetUserId(),
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "reward tracker not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}
	rewardTracker, err := collections_common.ToValue[collections_player.RewardTracker](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	//spin configure
	object, err = collections_system.ReadSpinConfigure(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "spin configure not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	spinConfigure, err := collections_common.ToValue[collections_system.SpinConfigure](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	now := time.Now()
	isSpinFree := spinConfigure.FreeSpinTime+rewardTracker.SpinInfo.LastSpinTime <= now.Unix()
	var nextFreeSpinCooldown int64
	if isSpinFree {
		nextFreeSpinCooldown = 0
	} else {
		// free time + last time - now
		nextFreeSpinCooldown = spinConfigure.FreeSpinTime + rewardTracker.SpinInfo.LastSpinTime - now.Unix()
	}

	claimedThisDay := rewardTracker.DailyRewardsInfo.LastClaimTime >= utils.StartOfToday(now).Unix()
	var nextDailyRewardCooldown int64
	if claimedThisDay {
		// start of next day - now
		nextDailyRewardCooldown = utils.StartOfTomorow(now).Unix() - now.Unix()
	} else {
		nextDailyRewardCooldown = 0
	}
	userCooldownTimers := UserCooldownTimers{
		NextFreeSpinCooldown:    nextFreeSpinCooldown,
		IsSpinFree:              isSpinFree,
		NextDailyRewardCooldown: nextDailyRewardCooldown,
		ClaimedThisDay:          claimedThisDay,
	}

	data, err := json.Marshal(userCooldownTimers)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = params.dispatcher.BroadcastMessage(OP_CODE_USER_COOLDOWN_TIMERS, data, []runtime.Presence{
		params.presence,
	}, nil, true)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
