package authenticator_graphql

import (
	cifarm_periphery "cifarm-server/src/services/periphery"
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"
	"github.com/heroiclabs/nakama-common/runtime"
)

type GetAuthenticationInput struct {
	AuthenticationId string `json:"authenticationId"`
}

type AuthenticationData struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
}

func GetAuthenticationData(
	ctx context.Context,
	logger runtime.Logger,
	input GetAuthenticationInput,
) (*AuthenticationData, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[cifarm_periphery.CIFARM_PERIPHERY_API_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_API_URL not found in environment variables")
		return nil, errors.New("CIFARM_PERIPHERY_API_URL not found in environment variables")
	}
	logger.Info("GRAPHPQL %v", url)
	client := graphql.NewClient(url, nil)

	query := `query Query($input: GetAuthenticationDataInput!) {
  authenticationData(input: $input) {
    chain,
	address
  }
}`
	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"authenticationId": input.AuthenticationId,
		},
	}
	result := struct {
		AuthenticationData AuthenticationData `json:"authenticationData"`
	}{}

	err := client.Exec(context.Background(),
		query,
		&result,
		variables,
	)
	if err != nil {
		return nil, err
	}
	return &result.AuthenticationData, nil
}
