package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftByTokenIdInput struct {
	TokenId  int    `json:"tokenId"`
	Network  string `json:"network"`
	NftKey   string `json:"nftKey"`
	ChainKey string `json:"chainKey"`
}

type GetNftByTokenIdArgs struct {
	Input GetNftByTokenIdInput `json:"input"`
}

func GetNftByTokenId(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	args GetNftByTokenIdArgs,
) (*NftDataResponse, error) {
	url, err := config.CifarmPeripheryGraphqlUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	client := graphql.NewClient(url, nil)
	logger.Info("%v", args.Input.TokenId)
	query := `query Query($args: GetNftByTokenIdArgs!) {
  nftByTokenId(args: $args) {
    ownerAddress,
    tokenId,
    tokenURI
  }
}`
	variables := map[string]interface{}{
		"args": args,
	}
	result := struct {
		NftByTokenId NftDataResponse `json:"nftByTokenId"`
	}{}

	err = client.WithDebug(true).Exec(context.Background(),
		query,
		&result,
		variables,
	)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	logger.Info(result.NftByTokenId.OwnerAddress)
	return &result.NftByTokenId, nil
}
