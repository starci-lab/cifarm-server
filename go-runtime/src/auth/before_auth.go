package auth

import (
	_constants "cifarm-server/src/constants"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

func BeforeAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	data *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {

	if data == nil {
		return nil, errors.New("data is nil")
	}

	chain := data.Account.Vars["chain"]

	_chain := _constants.DEFAULT_CHAIN

	if chain == "" {
		parsed, err := strconv.Atoi(chain)
		if err == nil {
			_chain = parsed
		}
	}

	value, err := json.Marshal(_collections.PlayerMetadataValue{
		Chain: _chain,
	})
	if err != nil {
		return nil, err
	}

	var userId string
	if !data.Create {
		userId = runtime.RUNTIME_CTX_USER_ID
	} else {
		userId =  
	}

	write := &runtime.StorageWrite{
		Collection:      _constants.CONFIG_COLLECTION,
		Key:             _constants.PLAYER_METADATA_KEY,
		UserID:          runtime.RUNTIME_CTX_USER_ID,
		Value:           string(value),
		PermissionRead:  2,
		PermissionWrite: 0,
	}

	writes := []*runtime.StorageWrite{
		write,
	}

	ack, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		return nil, err
	}
	logger.Info("Write ack: %v", ack)

	return data, nil
}
