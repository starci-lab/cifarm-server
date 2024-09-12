package rpcs_nfts

import (
	collections_nfts "cifarm-server/src/collections/nfts"
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

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

	var nfts []collections_nfts.Nft

	for _, nftResponse := range data.Records {
		nfts = append(nfts, collections_nfts.Nft{
			TokenId: nftResponse.TokenId,
			Type:    collections_nfts.TYPE_PREMIUM_TILE,
		})
	}
	err = collections_nfts.WriteMany(ctx, logger, db, nk, collections_nfts.WriteManyParams{
		Nfts:   nfts,
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(""), nil
}
