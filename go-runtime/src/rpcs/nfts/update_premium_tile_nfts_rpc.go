package rpcs_nfts

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_nfts "cifarm-server/src/collections/nfts"
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdatePremiumTileNftsRpcResponse struct {
	TokenIds []int `json:"tokenIds"`
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

	object, err := collections_config.ReadMetadataByKey(ctx, logger, db, nk, collections_config.ReadMetadataByKeyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	metadata, err := collections_common.ToValue[collections_config.Metadata](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	data, err := services_periphery_graphql.GetNfts(ctx, logger, services_periphery_graphql.GetNftArgs{
		Input: services_periphery_graphql.GetNftsInput{
			AccountAddress: metadata.AccountAddress,
			ChainKey:       metadata.ChainKey,
			Network:        metadata.Network,
			NftKey:         collections_nfts.NFT_KEY_PREMIUM_TILE,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var nfts []collections_nfts.Nft

	for _, nftResponse := range data.Records {
		nfts = append(nfts, collections_nfts.Nft{
			TokenId:        nftResponse.TokenId,
			Type:           collections_nfts.TYPE_PREMIUM_TILE,
			AccountAddress: metadata.AccountAddress,
			ChainKey:       metadata.ChainKey,
			Network:        metadata.Network,
		})
	}

	err = collections_nfts.WriteMany(ctx, logger, db, nk, collections_nfts.WriteManyParams{
		Nfts: nfts,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	return string(""), nil
}
