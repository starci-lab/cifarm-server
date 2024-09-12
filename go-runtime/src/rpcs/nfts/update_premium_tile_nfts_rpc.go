package rpcs_nfts

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_nfts "cifarm-server/src/collections/nfts"
	collections_tiles "cifarm-server/src/collections/tiles"
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"cifarm-server/src/utils"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteInventoriesAndDeleteFromOthersParams struct {
	ChainKey     string                 `json:"chainKey"`
	Network      string                 `json:"network"`
	TokenIds     []int                  `json:"tokenIds"`
	PreviousNfts []collections_nfts.Nft `json:"previousNfts"`
}

func WriteInventoriesThenDeleteFromOthers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteInventoriesAndDeleteFromOthersParams,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	objects, err := collections_nfts.ReadByTokenIds(ctx, logger, db, nk, collections_nfts.ReadByTokenIdsParams{
		TokenIds: params.TokenIds,
		ChainKey: params.ChainKey,
		Network:  params.Network,
		Type:     collections_nfts.TYPE_PREMIUM_TILE,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	nfts, err := collections_common.ToValues[collections_nfts.Nft](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	var newTokenIds []int
	//check if new tokens exist
	for _, nft := range nfts {
		//mean the nft is existed before in your wallet
		hasPreviousNft := false

		for _, previousNfts := range params.PreviousNfts {
			if nft.TokenId == previousNfts.TokenId {
				hasPreviousNft = true
				break
			}
		}
		//hasPreviousNft = false mean that you haven't have this before => new => create
		if !hasPreviousNft {
			//create inventory
			err := collections_inventories.Write(ctx, logger, db, nk, collections_inventories.WriteParams{
				Inventory: collections_inventories.Inventory{
					Key:          nft.MappingKey,
					ReferenceKey: collections_tiles.KEY_PREMIUM,
					Type:         collections_inventories.TYPE_TILE,
					TileInfo: collections_inventories.TileInfo{
						GrowthTimeReduction:  0,
						PestResistance:       0,
						ProductivityIncrease: 0,
						WeedResistance:       0,
					},
					Unique: true,
				},
				UserId: userId,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			//add to the list to track
			newTokenIds = append(newTokenIds, nft.TokenId)
		}
		//which mean maybe the nft is transfered from other address, it need to do a check
	}
	logger.Debug(utils.SliceToString(newTokenIds))
	//check if nonPreviousToken existed in other addresses, by calling an api

	return nil
}

func DeleteThenWriteInventoriesToOthers(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteInventoriesAndDeleteFromOthersParams,
) error {
	// 	//lose token, then process deletion
	// 	if !hasNft {
	// 		//get the nft key
	// 		object, err := collections_nfts.ReadByTokenId(ctx, logger, db, nk, collections_nfts.ReadByTokenIdParams{
	// 			TokenId:  nft.TokenId,
	// 			Type:     nft.Type,
	// 			ChainKey: nft.ChainKey,
	// 			Network:  nft.Network,
	// 		})
	// 		if err != nil {
	// 			logger.Error(err.Error())
	// 			return err
	// 		}
	// 		key := object.Key

	// 		//do deletion of yours
	// 		if nft.IsPlaced {
	// 			//delete placed items
	// 			err := collections_placed_items.Delete(ctx, logger, db, nk, collections_placed_items.DeleteParams{
	// 				Key:    nft.MappingKey,
	// 				UserId: userId,
	// 			})
	// 			if err != nil {
	// 				logger.Error(err.Error())
	// 				return err
	// 			}
	// 		} else {
	// 			//delete inventory
	// 			err := collections_inventories.Delete(ctx, logger, db, nk, collections_inventories.DeleteParams{
	// 				Key:      nft.MappingKey,
	// 				Quantity: 1,
	// 				UserId:   userId,
	// 			})
	// 			if err != nil {
	// 				logger.Error(err.Error())
	// 				return err
	// 			}
	// 		}
	// 	}
	// }
	return nil
}

type UpdatePremiumTileNftsRpcResponse struct {
	TokenIds []int `json:"tokenIds"`
}

func UpdatePremiumTileNftsRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return "", errors.New(errMsg)
	}

	object, err := collections_config.ReadMetadataByKey(ctx, logger, db, nk, collections_config.ReadMetadataByKeyParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	metadata, err := collections_common.ToValue[collections_config.Metadata](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	data, err := services_periphery_graphql.GetNfts(ctx, logger, services_periphery_graphql.GetNftArgs{
		Input: services_periphery_graphql.GetNftsInput{
			AccountAddress: metadata.AccountAddress,
			ChainKey:       metadata.ChainKey,
			Network:        metadata.Network,
			NftKey:         collections_nfts.NFT_KEY_PREMIUM_TILE,
		},
	})

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var nfts []collections_nfts.Nft

	for _, nftResponse := range data.Records {
		nfts = append(nfts, collections_nfts.Nft{
			TokenId:        nftResponse.TokenId,
			Type:           collections_nfts.TYPE_PREMIUM_TILE,
			AccountAddress: metadata.AccountAddress,
			ChainKey:       metadata.ChainKey,
			Network:        metadata.Network,
			MappingKey:     uuid.NewString(),
			IsPlaced:       false,
		})
	}

	err = collections_nfts.WriteMany(ctx, logger, db, nk, collections_nfts.WriteManyParams{
		Nfts: nfts,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	var tokenIds []int
	for _, nft := range nfts {
		tokenIds = append(tokenIds, nft.TokenId)
	}
	value, err := json.Marshal(UpdatePremiumTileNftsRpcResponse{
		TokenIds: tokenIds,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
