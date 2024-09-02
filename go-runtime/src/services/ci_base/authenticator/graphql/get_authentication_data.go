package authenticator_graphql

import (
	_constants "cifarm-server/src/constants"
	"context"
	"errors"

	"github.com/hasura/go-graphql-client"

	"github.com/heroiclabs/nakama-common/runtime"
)

type GetAuthenticationInput struct {
	AuthenticationId string `graphql:"authenticationId"`
}

type GetAuthenticationResult struct {
	Address string `graphql:"address"`
	Chain   string `graphql:"chain"`
}

func GetAuthenticationData(
	ctx context.Context,
	logger runtime.Logger,
	input GetAuthenticationInput,
) (*GetAuthenticationResult, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return nil, errors.New("cannot get environment variables")
	}
	url, ok := vars[_constants.ENV_CI_BASE_GRAPHQL_URL]
	if !ok {
		logger.Error("CI_BASE_GRAPHQL_URL not found in environment variables")
		return nil, errors.New("CI_GRAPHQL_API_URL not found in environment variables")
	}
	logger.Info("GRAPHPQL %v", url)
	client := graphql.NewClient(url, nil)

	variables := map[string]interface{}{
		"input": map[string]interface{}{
			"authenticationId": input.AuthenticationId,
		},
	}
	var q struct {
		Query struct {
			AuthenticationData GetAuthenticationResult `graphql:"authenticationData(input: $input)"`
		} `graphql:"Query($input: GetAuthenticationDataInput!)"`
	}
	err := client.Query(context.Background(), &q, variables)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &q.Query.AuthenticationData, nil
}
