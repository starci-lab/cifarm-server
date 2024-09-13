package services_periphery_graphql

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByOwnerAddressInput struct {
	AccountAddress string `json:"accountAddress"`
	Network        string `json:"network"`
	NftKey         string `json:"nftKey"`
	ChainKey       string `json:"chainKey"`
}

type GetNftsByOwnerAddressFilter struct {
	Skip int `json:"skip"`
	Take int `json:"take"`
}

type GetNftByOwnerAddressArgs struct {
	Input  GetNftsByOwnerAddressInput  `json:"input"`
	Filter GetNftsByOwnerAddressFilter `json:"filter"`
}

type GetNftsByOwnerAddressResponse struct {
	Records []NftDataResponse `json:"records"`
	Count   int               `json:"count"`
}

func GetNftsByOwnerAddress(
	ctx context.Context,
	logger runtime.Logger,
	args GetNftByOwnerAddressArgs,
) (*GetNftsByOwnerAddressResponse, error) {
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

	query := `query Query($args: GetNftsByOwnerAddressArgs!) {
  nftsByOwnerAddress(args: $args) {
    count,
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
		NftsByOwnerAddress GetNftsByOwnerAddressResponse `json:"nftsByOwnerAddress"`
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
	return &result.NftsByOwnerAddress, nil
}
