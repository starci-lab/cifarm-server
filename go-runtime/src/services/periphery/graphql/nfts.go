package services_periphery_graphql

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsInput struct {
	AccountAddress string `json:"accountAddress"`
	Network        string `json:"network"`
	NftKey         string `json:"nftKey"`
	ChainKey       string `json:"chainKey"`
}

type GetNftsFilter struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

type GetNftArgs struct {
	Input  GetNftsInput  `json:"input"`
	Filter GetNftsFilter `json:"filter"`
}

type NftResponse struct {
	TokenId  int    `json:"tokenId"`
	TokenURI string `json:"tokenURI"`
}

type GetNftsResponse struct {
	Records []NftResponse `json:"records"`
	Count   int           `json:"count"`
}

func GetNfts(
	ctx context.Context,
	logger runtime.Logger,
	args GetNftArgs,
) (*GetNftsResponse, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[cifarm_periphery.CIFARM_PERIPHERY_GRAPHQL_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_GRAPHQL_URL not found in environment variables")
		return nil, errors.New("CIFARM_PERIPHERY_GRAPHQL_URL not found in environment variables")
	}
	client := graphql.NewClient(url, nil)

	query := `query Query($args: GetNftsArgs!) {
  nfts(args: $args) {
    count,
    records {
      tokenId,
      tokenURI
    }
  }
}`
	variables := map[string]interface{}{
		"args": args,
	}
	result := struct {
		Nfts GetNftsResponse `json:"nfts"`
	}{}

	err := client.WithDebug(true).Exec(context.Background(),
		query,
		&result,
		variables,
	)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &result.Nfts, nil
}
