package auth

import (
	_constants "cifarm-server/src/constants"
	_inventories "cifarm-server/src/storage/inventories"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Claims struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func AfterAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	out *api.Session,
	in *api.AuthenticateCustomRequest,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	chain := in.Account.Vars["chain"]
	address := in.Account.Vars["address"]
	value, err := json.Marshal(_collections.PlayerMetadata{
		Chain:   chain,
		Address: address,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			UserID:          userId,
			Key:             _constants.KEY_PLAYER_METADATA,
			Collection:      _constants.COLLECTION_CONFIG,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = _inventories.WriteInventoryObject(ctx, logger, db, nk, _inventories.WriteInventoryObjectParams{
		Id:       _constants.FARMING_TILE_BASIC_FARMING_TILE_STARTER,
		Type:     _constants.TYPE_FARMING_TILE,
		Quantity: 6,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
