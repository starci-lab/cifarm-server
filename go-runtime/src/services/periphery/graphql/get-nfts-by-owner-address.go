package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByOwnerAddressInput struct {
	AccountAddress   string `json:"accountAddress"`
	Network          string `json:"network"`
	ChainKey         string `json:"chainKey"`
	NftCollectionKey string `json:"nftCollectionKey"`
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
	db *sql.DB,
	nk runtime.NakamaModule,
	args GetNftByOwnerAddressArgs,
) (*GetNftsByOwnerAddressResponse, error) {
	url, err := config.CifarmPeripheryGraphqlUrl(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	client := graphql.NewClient(url, nil)

	query := `query Query($args: GetNftsByOwnerAddressArgs!) {
  nftsByOwnerAddress(args: $args) {
    count,
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
		NftsByOwnerAddress GetNftsByOwnerAddressResponse `json:"nftsByOwnerAddress"`
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
	return &result.NftsByOwnerAddress, nil
}
