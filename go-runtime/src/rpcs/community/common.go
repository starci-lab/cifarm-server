package rpcs_community

import (
	"cifarm-server/src/friends"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type User struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type GetMutipleValueParams struct {
	UserId      string `json:"userId"`
	OtherUserId string `json:"otherUserId"`
}

func GetMutipleValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetMutipleValueParams,
) (int, error) {
	//check friend
	check, err := friends.CheckFriendByUserId(ctx, logger, db, nk, friends.CheckFriendByUserIdParams{
		UserId:       params.UserId,
		FriendUserId: params.OtherUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}
	multiplier := 1
	if check {
		multiplier = 2
	}
	return multiplier, nil
}
