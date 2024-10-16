package hooks

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func AfterAddFriends(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	in *api.AddFriendsRequest,
) error {
	// Handle the user's friends being added.
	logger.Info("AfterAddFriends")
	return nil
}
