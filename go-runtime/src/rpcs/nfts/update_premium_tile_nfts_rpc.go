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
	"sync"

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
	var wg sync.WaitGroup

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

	//users nft owned
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

	//previous users nft owned
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

	var tokenIds []int

	//WRITE: create new or received from other
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
			//nft not found in database => create new since new token is minted
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
			//nft found, mean that it is previously owned by other => do transfer ownership
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
	//DELETE: delete or transfer to other
	for _, previousNftInventory := range previousNftInventories {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			var found bool
			for _, nft := range data.Records {
				if nft.TokenId == previousNftInventory.TokenId {
					found = true
					break
				}
			}
			if !found {
				//not found mean previous nfts has the nft current do not, there are 2 case - burned or you transfer to others

				data, err := services_periphery_graphql.GetNftByTokenId(
					ctx,
					logger,
					services_periphery_graphql.GetNftByTokenIdArgs{
						Input: services_periphery_graphql.GetNftByTokenIdInput{
							TokenId:  previousNftInventory.TokenId,
							ChainKey: metadata.ChainKey,
							Network:  metadata.Network,
							NftKey:   "premiumTile",
						},
					})
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				if data == nil {
					//not found, mean it cannot be queried on chain => it have been burned
					err := collections_inventories.DeleteUnique(ctx, logger, db, nk, collections_inventories.DeleteUniqueParams{
						UserId: userId,
						Key:    previousNftInventory.Key,
					})
					if err != nil {
						logger.Error(err.Error())
						return err
					}
				} else {
					//the nft is still existed on chain, so that will 2 case
					newUserId, err := collections_config.GetUserIdByMetadata(ctx, logger, db, nk, collections_config.GetUserIdByMetadataParams{
						Metadata: collections_config.Metadata{
							ChainKey:       metadata.ChainKey,
							Network:        metadata.Network,
							AccountAddress: data.OwnerAddress,
						},
					})
					if err != nil {
						logger.Error(err.Error())
						return err
					}
					if newUserId != "" {
						//case 1, still existed in the database, so that we do transfer ownership
						err := collections_inventories.TransferOwnership(ctx, logger, db, nk, collections_inventories.TransferOwnershipParams{
							FromUserId: userId,
							ToUserId:   newUserId,
							Key:        previousNftInventory.Key,
						})
						if err != nil {
							logger.Error(err.Error())
							return err
						}
					} else {
						//nah, mean that it have been transfer to out-of-system address, just delete it
						err := collections_inventories.DeleteUnique(ctx, logger, db, nk, collections_inventories.DeleteUniqueParams{
							UserId: userId,
							Key:    previousNftInventory.Key,
						})
						if err != nil {
							logger.Error(err.Error())
							return err
						}
					}
				}
			}
			return nil
		}()
	}
	wg.Wait()

	value, err := json.Marshal(UpdatePremiumTileNftsRpcResponse{
		TokenIds: tokenIds,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
