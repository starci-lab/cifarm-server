package friends

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type CheckFriendByUsernameParams struct {
	UserId         string `json:"thatUserId"`
	FriendUsername string `json:"friendUsername"`
}

func CheckFriendByUsername(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params CheckFriendByUsernameParams,
) (bool, error) {
	var cursor string
	found := false
	for {
		friends, _, err := nk.FriendsList(ctx, params.UserId, 100, nil, cursor)
		if err != nil {
			logger.Error(err.Error())
			return false, err
		}

		for _, friend := range friends {
			if friend.User.Username == params.FriendUsername {
				found = true
				break
			}
		}

		if found {
			break
		}

		if len(friends) < 100 {
			break
		}
	}
	return found, nil
}

type CheckFriendByUserIdParams struct {
	UserId       string `json:"thatUserId"`
	FriendUserId string `json:"friendUserId"`
}

func CheckFriendByUserId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params CheckFriendByUserIdParams,
) (bool, error) {
	var cursor string
	found := false
	for {
		friends, _, err := nk.FriendsList(ctx, params.UserId, 100, nil, cursor)
		if err != nil {
			logger.Error(err.Error())
			return false, err
		}

		for _, friend := range friends {
			if friend.User.Id == params.UserId {
				found = true
				break
			}
		}

		if found {
			break
		}

		if len(friends) < 100 {
			break
		}
	}
	return found, nil
}
