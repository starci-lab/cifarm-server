package avalanche

import (
	"context"
	"database/sql"

	_common "cifarm-server/src/auth/common"
	_api "cifarm-server/src/utils/api"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func BeforeAvalancheAuth(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, data *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	endpoint := "/api/v1/verification"
	url := "https://blockchain-auth-service.starci.net/api" + endpoint

	response, err := _api.SendPostRequest[_common.VerifyMessageResponseData](url, data.Account.Id)
	if err != nil {
		logger.Error("Error sending POST request to blockchain auth server: %v", err)
		return nil, err
	}
	data.Account.Id = response.Address
	return data, err
}
