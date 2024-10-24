package rpcs_nfts

import (
	collections_common "cifarm-server/src/collections/common"
	collections_inventories "cifarm-server/src/collections/inventories"
	collections_player "cifarm-server/src/collections/player"
	collections_tiles "cifarm-server/src/collections/tiles"
	services_periphery_graphql "cifarm-server/src/services/periphery/graphql"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"sync"

	"github.com/heroiclabs/nakama-common/runtime"
)

type WriteOrTransferedFromParams struct {
	UserId string                                       `json:"userId"`
	Nfts   []services_periphery_graphql.NftDataResponse `json:"nfts"`
}

func WriteOrTransferedFrom(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params WriteOrTransferedFromParams,
) ([]string, error) {
	var tokenIds []string
	for _, nftResponse := range params.Nfts {
		object, err := collections_inventories.ReadByTokenId(ctx, logger, db, nk, collections_inventories.ReadByTokenIdParams{
			TokenId:      nftResponse.TokenId,
			ReferenceKey: collections_tiles.KEY_FERTILE_TILE,
		})
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		if object == nil {
			//nft not found in database => create new since new token is minted
			logger.Debug("Case 1.1: Nft not found in datate, so that we create new: %v", nftResponse.TokenId)

			_, err := collections_inventories.WriteUnique(ctx, logger, db, nk, collections_inventories.WriteUniqueParams{
				UserId: params.UserId,
				Inventory: collections_inventories.Inventory{
					ReferenceKey: collections_tiles.KEY_FERTILE_TILE,
					Placeable:    true,
					TokenId:      nftResponse.TokenId,
					Type:         collections_inventories.TYPE_TILE,
				},
			})
			if err != nil {
				logger.Error(err.Error())
				return nil, err
			}
		} else {
			//nft found, mean that it is previously owned by other => do transfer ownership
			logger.Debug("Case 1.2: Nft found in database, mean that it is previously owned by other, do transfer ownership: %v", nftResponse.TokenId)

			if object.UserId != params.UserId {
				err := collections_inventories.TransferOwnership(ctx, logger, db, nk, collections_inventories.TransferOwnershipParams{
					FromUserId: object.UserId,
					ToUserId:   params.UserId,
					Key:        object.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return nil, err
				}
			}
		}
		tokenIds = append(tokenIds, nftResponse.TokenId)
	}
	return tokenIds, nil
}

type DeleteOrTransferToParams struct {
	UserId                 string                                       `json:"userId"`
	Metadata               *collections_player.Metadata                 `json:"metadata"`
	Nfts                   []services_periphery_graphql.NftDataResponse `json:"nfts"`
	PreviousNftInventories []*collections_inventories.Inventory         `json:"previousNftInventories"`
}

func DeleteOrTransferTo(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params DeleteOrTransferToParams,
) error {
	var wg sync.WaitGroup
	for _, previousNftInventory := range params.PreviousNftInventories {
		wg.Add(1)
		go func() error {
			defer wg.Done()
			var found bool
			for _, nft := range params.Nfts {
				if nft.TokenId == previousNftInventory.TokenId {
					found = true
					break
				}
			}
			if found {
				return nil
			}
			//not found mean previous nfts has the nft current do not, there are 2 case - burned or you transfer to others

			data, err := services_periphery_graphql.GetNftByTokenId(
				ctx,
				logger,
				db,
				nk,
				services_periphery_graphql.GetNftByTokenIdArgs{
					Input: services_periphery_graphql.GetNftByTokenIdInput{
						TokenId:          previousNftInventory.TokenId,
						ChainKey:         params.Metadata.ChainKey,
						Network:          params.Metadata.Network,
						NftCollectionKey: collections_inventories.NFT_FERTILE_TILE,
					},
				})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
			if data == nil {
				//not found, mean it cannot be queried on chain => it have been burned
				logger.Debug("Case 2.1: Cannot be queried on chain: %v", previousNftInventory.TokenId)

				err := collections_inventories.DeleteUnique(ctx, logger, db, nk, collections_inventories.DeleteUniqueParams{
					UserId: params.UserId,
					Key:    previousNftInventory.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return err
				}

				return nil
			}

			//the nft is still existed on chain, so that will 2 case
			logger.Debug("Case 2.2: Existed on chain, but you already transfer it: %v", previousNftInventory.TokenId)

			newUserId, err := collections_player.GetUserIdByMetadata(ctx, logger, db, nk, collections_player.GetUserIdByMetadataParams{
				Metadata: collections_player.Metadata{
					ChainKey:       params.Metadata.ChainKey,
					Network:        params.Metadata.Network,
					AccountAddress: data.OwnerAddress,
				},
			})

			if err != nil {
				logger.Error(err.Error())
				return err
			}
			if newUserId != "" {
				//case 1, still existed in the database, so that we do transfer ownership
				logger.Debug("Case 2.2.1: The nft is still in the database, so that we do transfer ownership: %v", previousNftInventory.TokenId)

				err := collections_inventories.TransferOwnership(ctx, logger, db, nk, collections_inventories.TransferOwnershipParams{
					FromUserId: params.UserId,
					ToUserId:   newUserId,
					Key:        previousNftInventory.Key,
				})
				if err != nil {
					logger.Error(err.Error())
					return err
				}
				return nil
			}
			//nah, mean that it have been transfer to out-of-system address, just delete it
			logger.Debug("Case 2.2.2: The nft has transfered to out-of-system address, just delete it: %v", previousNftInventory.TokenId)

			err = collections_inventories.DeleteUnique(ctx, logger, db, nk, collections_inventories.DeleteUniqueParams{
				UserId: params.UserId,
				Key:    previousNftInventory.Key,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			return nil
		}()
	}
	wg.Wait()
	return nil
}

type UpdateFertileTileNftsRpcResponse struct {
	TokenIds []string `json:"tokenIds"`
}

func UpdateFertileTileNftsRpc(
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

	object, err := collections_player.ReadMetadata(ctx, logger, db, nk, collections_player.ReadMetadataParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	metadata, err := collections_common.ToValue[collections_player.Metadata](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//users nft owned
	data, err := services_periphery_graphql.GetNftsByOwnerAddress(
		ctx,
		logger,
		db,
		nk,
		services_periphery_graphql.GetNftByOwnerAddressArgs{
			Input: services_periphery_graphql.GetNftsByOwnerAddressInput{
				AccountAddress:   metadata.AccountAddress,
				ChainKey:         metadata.ChainKey,
				Network:          metadata.Network,
				NftCollectionKey: collections_inventories.NFT_FERTILE_TILE,
			},
		})

	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//previous users nft owned
	objects, err := collections_inventories.ReadManyUnique(ctx, logger, db, nk, collections_inventories.ReadManyUniqueParams{
		UserId:       userId,
		ReferenceKey: collections_tiles.KEY_FERTILE_TILE,
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

	//WRITE: create new or received from other
	tokenIds, err := WriteOrTransferedFrom(ctx, logger, db, nk, WriteOrTransferedFromParams{
		UserId: userId,
		Nfts:   data.Records,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//DELETE: delete or transfer to other
	err = DeleteOrTransferTo(ctx, logger, db, nk, DeleteOrTransferToParams{
		UserId:                 userId,
		Metadata:               metadata,
		Nfts:                   data.Records,
		PreviousNftInventories: previousNftInventories,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	value, err := json.Marshal(UpdateFertileTileNftsRpcResponse{
		TokenIds: tokenIds,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return string(value), nil
}
