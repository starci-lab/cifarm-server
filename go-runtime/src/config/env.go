package config

import (
	"context"
	"database/sql"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	CIFARM_PERIPHERY_API_URL     = "cifarm-periphery-api-url"
	CIFARM_PERIPHERY_GRAPHQL_URL = "cifarm-periphery-graphql-url"
	MINTER_PRIVATE_KEY           = "minter-private-key"
	UTILITY_TOKEN_ADDRESS        = "utility-token-address"
	AUTHENTICATION_ID            = "authentication-id"
)

func MinterPrivateKey(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (string, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return "", errors.New("cannot get environment variables")
	}
	//get env
	value, ok := vars[MINTER_PRIVATE_KEY]
	if !ok {
		logger.Error("MINTER_PRIVATE_KEY not found in environment variables")
		return "", errors.New("MINTER_PRIVATE_KEY not found in environment variables")
	}
	return value, nil
}

func UtilityTokenAddress(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (string, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return "", errors.New("cannot get environment variables")
	}
	//get env
	value, ok := vars[UTILITY_TOKEN_ADDRESS]
	if !ok {
		logger.Error("UTILITY_TOKEN_ADDRESS not found in environment variables")
		return "", errors.New("UTILITY_TOKEN_ADDRESS not found in environment variables")
	}
	return value, nil
}

func CifarmPeripheryApiUrl(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (string, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return "", errors.New("cannot get environment variables")
	}
	//get env
	value, ok := vars[CIFARM_PERIPHERY_API_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_API_URL not found in environment variables")
		return "", errors.New("CIFARM_PERIPHERY_API_URL not found in environment variables")
	}
	return value, nil
}

func CifarmPeripheryGraphqlUrl(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (string, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return "", errors.New("cannot get environment variables")
	}
	//get env
	value, ok := vars[CIFARM_PERIPHERY_GRAPHQL_URL]
	if !ok {
		logger.Error("CIFARM_PERIPHERY_GRAPHQL_URL not found in environment variables")
		return "", errors.New("CIFARM_PERIPHERY_GRAPHQL_URL not found in environment variables")
	}
	return value, nil
}

func AuthenticationId(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (string, error) {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		logger.Error("Cannot get environment variables")
		return "", errors.New("cannot get environment variables")
	}
	//get env
	value, ok := vars[AUTHENTICATION_ID]
	if !ok {
		logger.Error("AUTHENTICATION_ID not found in environment variables")
		return "", errors.New("AUTHENTICATION_ID not found in environment variables")
	}
	return value, nil
}
