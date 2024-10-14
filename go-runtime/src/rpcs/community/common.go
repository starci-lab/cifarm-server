package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/friends"
	"context"
	"database/sql"
	"math"
	"math/rand"

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

type GetThiefValueParams struct {
	MaximunTheifQuantity int `json:"maximunTheifQuantity"`
}

func GetThiefValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetThiefValueParams,
) (int, error) {
	object, err := collections_system.ReadGlobalConstants(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}
	if object == nil {
		errMsg := "global constants not found"
		logger.Error(errMsg)
		return 0, err
	}
	globalConstants, err := collections_common.ToValue[collections_system.GlobalConstants](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	thiefQuantity := 1
	random := rand.Float64()
	if random > globalConstants.GameRandomness.Theif3 {
		thiefQuantity = 3
	} else if random > globalConstants.GameRandomness.Theif2 {
		thiefQuantity = 2
	}
	thiefQuantity = int(math.Min(float64(params.MaximunTheifQuantity), float64(thiefQuantity)))
	return thiefQuantity, nil
}
