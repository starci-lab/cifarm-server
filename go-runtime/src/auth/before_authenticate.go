package auth

import (
	services_cibase_authenticator_api "cifarm-server/src/services/cibase/api/authenticator"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func BeforeAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	data *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {

	message, ok := data.Account.Vars["message"]
	if !ok {
		return nil, errors.New("missing 'message' in account variables")
	}

	signature, ok := data.Account.Vars["signature"]
	if !ok {
		return nil, errors.New("missing 'signature' in account variables")
	}

	publicKey, ok := data.Account.Vars["publicKey"]
	if !ok {
		return nil, errors.New("missing 'publicKey' in account variables")
	}

	chain, ok := data.Account.Vars["chain"]
	if !ok {
		return nil, errors.New("missing 'chain' in account variables")
	}

	body := services_cibase_authenticator_api.VerifyMessageRequestBody{
		Message:   message,
		PublicKey: publicKey,
		Signature: signature,
		Chain:     chain,
	}

	response, err := services_cibase_authenticator_api.VerifyMessage(ctx, logger, &body)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	data.Account.Id = response.AuthenticationId
	data.Username = fmt.Sprintf("%s_%s", chain, response.Address)
	data.Create.Value = true
	data.Account.Vars["address"] = response.Address
	return data, nil
}
