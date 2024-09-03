package auth

import (
	_constants "cifarm-server/src/constants"
	_authenticator_graphql "cifarm-server/src/services/ci_base/authenticator/graphql"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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

	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return fmt.Errorf("user ID not found")
	}

	value, err := json.Marshal(_collections.PlayerMetadata{
		Chain:   response.Chain,
		Address: response.Address,
	})
	if err != nil {
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			UserID:          userId,
			Key:             _constants.KEY_PLAYER_METADATA,
			Collection:      _constants.COLLECTION_CONFIG,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	err = nk.AccountUpdateId(
		ctx,
		userId, "",
		nil,
		fmt.Sprintf("%s_%s", response.Chain, response.Address),
		"", "", "", "")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
