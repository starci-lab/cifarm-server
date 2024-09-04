package entities

import (
	_constants "cifarm-server/src/constants"
	_lands "cifarm-server/src/storage/lands"
	_collections "cifarm-server/src/types/collections"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func SetupLand(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {

	lands := []_collections.Land{
		{
			Id:           _constants.LAND_BASIC_LAND,
			InitialPrice: 1000,
		},
		{
			Id: _constants.LAND_PREMIUM_LAND,
		},
	}

	err := _lands.WriteLandObjects(ctx, logger, db, nk, _lands.WriteLandObjectsParams{
		Lands: lands,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
