package avalanche

import (
	"context"
	"database/sql"
	"errors"

	_common "cifarm-server/src/auth/common"
	_constants "cifarm-server/src/constants"
	_api "cifarm-server/src/utils/api"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func BeforeAvalancheAuth(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, data *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[_constants.ENV_BLOCKCHAIN_AUTH_SERVER_URL]
	if !ok {
		logger.Error("Error getting blockchain auth server URL: %v", url)
		return nil, errors.New("error getting blockchain auth server URL")
	} else {
		logger.Error("Error getting blockchain auth server URL: %v", url)
	}
	endpoint := "/v1/verifications"
	url = url + endpoint

	response, err := _api.SendPostRequest[_common.VerifyMessageResponseData](url, data.Account.Id)
	if err != nil {
		logger.Error("Error sending POST request to blockchain auth server: %v", err)
		return nil, err
	}
	data.Account.Id = response.Address
	return data, err
}
