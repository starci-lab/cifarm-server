package rpcs_nfts

import (
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdatePremiumTileNftsRpcParams struct {
	AccountAddress string `json:"accountAddress"`
}

func UpdatePremiumTileNftsRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	var params *UpdatePremiumTileNftsRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	data, err := services_periphery_graphql.GetNfts(ctx, logger, services_periphery_graphql.GetNftArgs{
		Input: services_periphery_graphql.GetNftsInput{
			AccountAddress: params.AccountAddress,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	logger.Info("%v", data.Count)

	return string(""), nil
}
