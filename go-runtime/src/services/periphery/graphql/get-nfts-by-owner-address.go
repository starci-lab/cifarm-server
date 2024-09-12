package services_periphery_graphql

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByTokenIdsInput struct {
	TokenIds []int  `json:"tokenIds"`
	Network  string `json:"network"`
	NftKey   string `json:"nftKey"`
	ChainKey string `json:"chainKey"`
}

type GetNftsByTokenIdsArgs struct {
	Input GetNftsByTokenIdsInput `json:"input"`
}

type GetNftsByTokenIdsResponse struct {
	Records []NftData `json:"records"`
	Count   int       `json:"count"`
}

func GetNftsByTokenIds(
	ctx context.Context,
	logger runtime.Logger,
	args GetNftsByTokenIdsArgs,
) (*GetNftsByTokenIdsResponse, error) {
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

	query := `query Query($args: GetNftsByTokenIdsArgs!) {
  nftsByTokenIds(args: $args) {
    records {
      tokenId,
      tokenURI,
	  ownerAddress
    }
  }
}`
	variables := map[string]interface{}{
		"args": args,
	}
	result := struct {
		NftsByTokenIds GetNftsByTokenIdsResponse `json:"nftsByTokenIds"`
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
	return &result.NftsByTokenIds, nil
}
