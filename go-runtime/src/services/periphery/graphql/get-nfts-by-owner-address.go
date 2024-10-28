package services_periphery_graphql

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetNftsByOwnerAddressInput struct {
	AccountAddress   string `json:"accountAddress,omitempty"`
	Network          string `json:"network,omitempty"`
	ChainKey         string `json:"chainKey,omitempty"`
	NftCollectionKey string `json:"nftCollectionKey,omitempty"`
}

type GetNftsByOwnerAddressFilter struct {
	Skip int `json:"skip,omitempty"`
	Take int `json:"take,omitempty"`
}

type GetNftByOwnerAddressArgs struct {
	Input  GetNftsByOwnerAddressInput  `json:"input,omitempty"`
	Filter GetNftsByOwnerAddressFilter `json:"filter,omitempty"`
}

type GetNftsByOwnerAddressResponse struct {
	Records []NftDataResponse `json:"records,omitempty"`
	Count   int               `json:"count,omitempty"`
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
		NftsByOwnerAddress GetNftsByOwnerAddressResponse `json:"nftsByOwnerAddress,omitempty"`
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
