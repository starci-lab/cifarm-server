package collections_nfts

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteParams struct {
	Nft    Nft    `json:"nft"`
	UserId string `json:"userId"`
}

func Write(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteParams,
) error {
	key := uuid.NewString()
	params.Nft.Key = key

	data, err := json.Marshal(
		params.Nft,
	)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection:      COLLECTION_NAME,
			Key:             key,
			UserID:          params.UserId,
			Value:           string(data),
			PermissionRead:  2,
			PermissionWrite: 0,
		},
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

type WriteManyParams struct {
	Nfts   []Nft  `json:"nfts"`
	UserId string `json:"userId"`
}

func WriteMany(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteManyParams,
) error {
	var writes []*runtime.StorageWrite
	for _, nft := range params.Nfts {
		//check existed, then override
		foundNft, err := ReadByTokenId(ctx, logger, db, nk, ReadByTokenIdParams{
			TokenId: nft.TokenId,
			Type:    nft.Type,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		if foundNft == nil {
			nft.Key = uuid.NewString()
		} else {
			nft.Key = foundNft.Key
		}

		value, err := json.Marshal(nft)
		if err != nil {
			continue
		}

		write := &runtime.StorageWrite{
			Collection:      COLLECTION_NAME,
			Key:             nft.Key,
			UserID:          params.UserId,
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
