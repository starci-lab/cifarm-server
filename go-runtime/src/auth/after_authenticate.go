package auth

import (
	_authenticator_graphql "cifarm-server/src/services/ci_base/authenticator/graphql"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Claims struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func AfterAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	out *api.Session,
	in *api.AuthenticateCustomRequest,
) error {
	input := _authenticator_graphql.GetAuthenticationInput{
		AuthenticationId: in.Account.Id,
	}
	response, err := _authenticator_graphql.GetAuthenticationData(ctx, logger, input)
	if err != nil {
		return err
	}
	logger.Info("aaa %s %s", response.Address, response.Chain)
	return nil
}
