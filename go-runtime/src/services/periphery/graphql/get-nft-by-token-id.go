package services_periphery_graphql

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	"context"
	"errors"

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
	args GetNftByTokenIdArgs,
) (*NftDataResponse, error) {
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

	err := client.WithDebug(true).Exec(context.Background(),
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
