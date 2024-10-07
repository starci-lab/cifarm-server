package auth

import (
	services_periphery_api_authenticator "cifarm-server/src/services/periphery/api/authenticator"
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

	chainKey, ok := data.Account.Vars["chainKey"]
	if !ok {
		return nil, errors.New("missing 'chainKey' in account variables")
	}

	network, ok := data.Account.Vars["network"]
	if !ok {
		return nil, errors.New("missing 'network' in account variables")
	}

	body := services_periphery_api_authenticator.VerifyMessageRequestBody{
		Message:   message,
		PublicKey: publicKey,
		Signature: signature,
		ChainKey:  chainKey,
		Network:   network,
	}

	response, err := services_periphery_api_authenticator.VerifyMessage(ctx, logger, db, nk, services_periphery_api_authenticator.VerifyMessageParams{
		Body: body,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	telegramInitDataRaw, ok := data.Account.Vars["telegramInitDataRaw"]
	if !ok {
		return nil, errors.New("missing 'telegramInitDataRaw' in account variables")
	}

	authorizeTelegramResponse, err := services_periphery_api_authenticator.AuthorizeTelegram(ctx, logger, db, nk, services_periphery_api_authenticator.AuthorizeTelegramParams{
		TelegramInitDataRaw: telegramInitDataRaw,
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	data.Account.Id = response.AuthenticationId
	data.Create.Value = true
	data.Username = fmt.Sprintf("%s_%s", chainKey, response.Address)
	data.Account.Vars["accountAddress"] = response.Address
	data.Account.Vars["telegramUserId"] = authorizeTelegramResponse.TelegramData.UserId
	return data, nil
}
