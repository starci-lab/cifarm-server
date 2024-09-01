package avalanche

import (
	_blockchain_auth_service "cifarm-server/src/services/blockchain_auth_service"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func BeforeAvalancheAuth(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	data *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {

	if data == nil {
		return nil, errors.New("data is nil")
	}

	message := data.Account.Vars["message"]
	signature := data.Account.Vars["signature"]
	publicKey := data.Account.Vars["publicKey"]
	platform := data.Account.Vars["platform"]

	var _platform *string = &platform
	if platform == "" {
		_platform = nil
	}

	body := _blockchain_auth_service.VerifyMessageRequestBody{
		Message:   message,
		Signature: signature,
		PublicKey: publicKey,
		Platform:  _platform,
	}

	response, err := _blockchain_auth_service.VerifyMessage(ctx, logger, &body)
	data.Account.Id = response.Address
	return data, err
}
