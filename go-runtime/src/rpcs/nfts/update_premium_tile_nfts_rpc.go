package rpcs_nfts

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_tiles "cifarm-server/src/collections/tiles"
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/heroiclabs/nakama-common/runtime"
)

type UpdatePremiumTileNftsRpcResponse struct {
	TokenIds []int `json:"tokenIds"`
}

const (
	PREMIUM_TILE = "premiumTile"
)

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

	data, err := services_periphery_graphql.GetNftsByOwnerAddress(
		ctx,
		logger,
		services_periphery_graphql.GetNftByOwnerAddressArgs{
			Input: services_periphery_graphql.GetNftsByOwnerAddressInput{
				AccountAddress: metadata.AccountAddress,
				ChainKey:       metadata.ChainKey,
				Network:        metadata.Network,
				NftKey:         "premiumTile",
			},
		})

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//previousNft
	objects, err := collections_inventories.ReadManyByUserId(ctx, logger, db, nk, collections_inventories.ReadManyByUserIdParams{
		UserId:       userId,
		ReferenceKey: collections_tiles.KEY_PREMIUM,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	previousNftInventories, err := collections_common.ToValues[collections_inventories.Inventory](ctx, logger, db, nk, objects)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//create or transfer from other
	var tokenIds []int
	for _, nftResponse := range data.Records {
		object, err := collections_inventories.ReadByTokenId(ctx, logger, db, nk, collections_inventories.ReadByTokenIdParams{
			TokenId:      nftResponse.TokenId,
			ReferenceKey: collections_tiles.KEY_PREMIUM,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
		if object == nil {
			//nft not found, mean that you have create new
			err := collections_inventories.WriteUnique(ctx, logger, db, nk, collections_inventories.WriteUniqueParams{
				UserId: userId,
				Inventory: collections_inventories.Inventory{
					ReferenceKey: collections_tiles.KEY_PREMIUM,
					Placeable:    true,
					TokenId:      nftResponse.TokenId,
					Type:         collections_inventories.TYPE_TILE,
				},
				PermissionRead: 2,
			})
			if err != nil {
				logger.Error(err.Error())
				return "", err
			}
		} else {
			//nft found, then we check if previous owner difference from now, so that we procedure a transfer ownership
			if object.UserId != userId {
				err := collections_inventories.TransferOwnership(ctx, logger, db, nk, collections_inventories.TransferOwnershipParams{
					FromUserId: object.UserId,
					ToUserId:   userId,
					Key:        object.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return "", err
				}
			}
		}
		tokenIds = append(tokenIds, nftResponse.TokenId)
	}
	//DESTROY or you transfer to others
	for _, previousNftInventory := range previousNftInventories {
		var found bool
		for _, nft := range data.Records {
			if nft.TokenId == previousNftInventory.TokenId {
				found = true
				break
				//do nothing
			}
		}
		if !found {
			//not found mean that the previous nft have been disable, there is too case - lose or you transfer
			object, err := collections_inventories.ReadByTokenId(ctx, logger, db, nk, collections_inventories.ReadByTokenIdParams{
				TokenId:      previousNftInventory.TokenId,
				ReferenceKey: collections_tiles.KEY_PREMIUM,
			})
			if err != nil {
				logger.Error(err.Error())
				return "", err
			}
			if object == nil {
				//nft not found, mean that the token has been removed
				err := collections_inventories.DeleteUnique(ctx, logger, db, nk, collections_inventories.DeleteUniqueParams{
					UserId: userId,
					Key:    previousNftInventory.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return "", err
				}
			} else {
				//you must transfer ownership
				err := collections_inventories.TransferOwnership(ctx, logger, db, nk, collections_inventories.TransferOwnershipParams{
					FromUserId: userId,
					ToUserId:   object.UserId,
					Key:        object.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return "", err
				}
			}
		}
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
