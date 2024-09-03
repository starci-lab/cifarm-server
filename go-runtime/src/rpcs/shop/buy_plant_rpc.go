package shop

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type BuyPlantRpcParams struct {
	Id       bool `json:"id"`
	Quantity int  `json:"quantity"`
}

func BuyPlantRpc(ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string) (string, error) {

	_, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	var params *BuyPlantRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	// plants, err := nk.StorageRead(ctx, []*runtime.StorageRead{
	// 	{
	// 		Collection: _constants.COLLECTION_ENTITIES,
	// 		Key:        _constants.KEY_PLANTS,
	// 	},
	// })
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return "", err
	// }
	// if len(plants) == 0 {
	// 	errMsg := "no plant found"
	// 	logger.Error(errMsg)
	// 	return "", errors.New(errMsg)
	// }

	// plant := plants[0]
	// var _plants _collections.Plants
	// err = json.Unmarshal([]byte(plant.Value), &_plants)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return "", err
	// }

	// for _, item := range _plants.Items {
	// 	if item.Id == params.Id {
	// 		return &item, true
	// 	}
	// }
	// price:=
	return "", nil
}
