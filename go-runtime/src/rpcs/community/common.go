package rpcs_community

import (
	collections_common "cifarm-server/src/collections/common"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"math"
	"math/rand"

	"github.com/heroiclabs/nakama-common/runtime"
)

type User struct {
	UserId   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
}

type GetThiefValueParams struct {
	MaximunTheifQuantity int `json:"maximunTheifQuantity,omitempty"`
}

func GetThiefValue(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params GetThiefValueParams,
) (int, error) {
	object, err := collections_system.ReadCropRandomness(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}
	if object == nil {
		errMsg := "global constants not found"
		logger.Error(errMsg)
		return 0, err
	}
	cropRandomness, err := collections_common.ToValue[collections_system.CropRandomness](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}

	thiefQuantity := 1
	random := rand.Float64()
	if random > cropRandomness.Theif3 {
		thiefQuantity = 3
	} else if random > cropRandomness.Theif2 {
		thiefQuantity = 2
	}
	thiefQuantity = int(math.Min(float64(params.MaximunTheifQuantity), float64(thiefQuantity)))
	return thiefQuantity, nil
}
