package collections_market_pricings

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteManyParams struct {
	MarketPricings []MarketPricing `json:"marketPricings"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, marketPricing := range params.MarketPricings {
		key := marketPricing.Key
		marketPricing.Key = ""
		value, err := json.Marshal(marketPricing)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      COLLECTION_NAME,
			Key:             key,
			Value:           string(value),
			PermissionRead:  2,
			PermissionWrite: 0,
		}
		writes = append(writes, write)
	}

	_, err := nk.StorageWrite(ctx, writes)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
