package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByTokenIdsInput struct {
	TokenIds         []string `json:"tokenIds,omitempty"`
	Network          string   `json:"network,omitempty"`
	NftCollectionKey string   `json:"nftCollectionKey,omitempty"`
	ChainKey         string   `json:"chainKey,omitempty"`
}

type GetNftsByTokenIdsArgs struct {
	Input GetNftsByTokenIdsInput `json:"input,omitempty"`
}

type GetNftsByTokenIdsResponse struct {
	Records []NftDataResponse `json:"records,omitempty"`
	Count   int               `json:"count,omitempty"`
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
		NftsByTokenIds GetNftsByTokenIdsResponse `json:"nftsByTokenIds,omitempty"`
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
