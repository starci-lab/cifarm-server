package rpcs_users

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Initialize(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	err := initializer.RegisterRpc("check_friend_by_user_id", CheckFriendByUserIdRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("check_friend_by_username", CheckFriendByUsernameRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("visit_rpc", VisitRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("thief_plant", ThiefPlantRpc)
	if err != nil {
		return err
	}

	err = initializer.RegisterRpc("search_users", SearchUserRpc)
	if err != nil {
		return err
	}

	return nil
}
