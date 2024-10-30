package hooks

import (
	services_periphery_api_authenticator "cifarm-server/src/services/periphery/api/authenticator"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
		errMsg := "missing 'message' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	signature, ok := data.Account.Vars["signature"]
	if !ok {
		errMsg := "missing 'signature' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	publicKey, ok := data.Account.Vars["publicKey"]
	if !ok {
		errMsg := "missing 'publicKey' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	chainKey, ok := data.Account.Vars["chainKey"]
	if !ok {
		errMsg := "missing 'chainKey' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	network, ok := data.Account.Vars["network"]
	if !ok {
		errMsg := "missing 'network' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	accountAddress, ok := data.Account.Vars["accountAddress"]
	if !ok {
		errMsg := "missing 'accountAddress' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	telegramInitDataRaw, ok := data.Account.Vars["telegramInitDataRaw"]
	if !ok {
		errMsg := "missing 'telegramInitDataRaw' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
	}

	botType, ok := data.Account.Vars["botType"]
	if !ok {
		errMsg := "missing 'botType' in account variables"
		logger.Error(errMsg)
		return nil, errors.New(errMsg)
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
		return nil, err
	}

	authorizeTelegramResponse, err := services_periphery_api_authenticator.AuthorizeTelegram(ctx, logger, db, nk, services_periphery_api_authenticator.AuthorizeTelegramParams{
		TelegramInitDataRaw: telegramInitDataRaw,
		BotType:             botType,
	})

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	data.Account.Id = response.AuthenticationId
	data.Create.Value = true
	data.Username = fmt.Sprintf("%s_%s", chainKey, accountAddress)

	_userId := strconv.Itoa(authorizeTelegramResponse.TelegramData.UserId)
	data.Account.Vars["telegramUserId"] = _userId
	return data, nil
}
