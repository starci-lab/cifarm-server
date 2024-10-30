package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByTokenIdsInput struct {
	TokenIds         []string `json:"tokenIds"`
	Network          string   `json:"network"`
	NftCollectionKey string   `json:"nftCollectionKey"`
	ChainKey         string   `json:"chainKey"`
}

type GetNftsByTokenIdsArgs struct {
	Input GetNftsByTokenIdsInput `json:"input"`
}

type GetNftsByTokenIdsResponse struct {
	Records []NftDataResponse `json:"records"`
	Count   int               `json:"count"`
}

func GetNftsByTokenIds(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	args GetNftsByTokenIdsArgs,
) (*GetNftsByTokenIdsResponse, error) {
	url, err := config.CifarmPeripheryGraphqlUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	client := graphql.NewClient(url, nil)

	query := `query Query($args: GetNftsByTokenIdsArgs!) {
  nftsByTokenIds(args: $args) {
    records {
      tokenId,
	  ownerAddress,
	  metadata {
        image,
        properties
      }
    }
  }
}`
	variables := map[string]interface{}{
		"args": args,
	}
	result := struct {
		NftsByTokenIds GetNftsByTokenIdsResponse `json:"nftsByTokenIds"`
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
	return &result.NftsByTokenIds, nil
}
