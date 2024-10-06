package rpcs_internal

import (
	collections_common "cifarm-server/src/collections/common"
	collections_config "cifarm-server/src/collections/config"
	collections_system "cifarm-server/src/collections/system"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type InternalInviteRpcParams struct {
	UserId       string `json:"userId"`
	OriginUserId string `json:"originUserId"`
}

func InternalInviteRpc(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	payload string,
) (string, error) {
	//permission check
	ok := CheckPermission(ctx, logger, db, nk)
	if !ok {
		return "", nil
	}

	var params *InternalInviteRpcParams
	err := json.Unmarshal([]byte(payload), &params)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	//invite rpc
	// 1, 3, 10, 25
	object, err := collections_config.ReadPlayerStats(ctx, logger, db, nk, collections_config.ReadPlayerStatsParams{
		UserId: params.OriginUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return "", nil
	}
	playerStats, err := collections_common.ToValue[collections_config.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	playerStats.Invites = append(playerStats.Invites, params.OriginUserId)

	err = collections_config.WritePlayerStats(ctx, logger, db, nk, collections_config.WritePlayerStatsParams{
		PlayerStats: *playerStats,
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	object, err = collections_system.ReadRewards(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}
	if object == nil {
		errMsg := "rewards not found"
		logger.Error(errMsg)
		return "", nil
	}
	rewards, err := collections_common.ToValue[collections_system.Rewards](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	if len(playerStats.Invites) == rewards.FromInvites.Metrics[1].Key {
		err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
			UserId: params.OriginUserId,
			Amount: rewards.FromInvites.Metrics[1].Value,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[2].Key {
		err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
			UserId: params.OriginUserId,
			Amount: rewards.FromInvites.Metrics[2].Value,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[3].Key {
		err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
			UserId: params.OriginUserId,
			Amount: rewards.FromInvites.Metrics[3].Value,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}

	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[4].Key {
		err = wallets.UpdateWalletGolds(ctx, logger, db, nk, wallets.UpdateWalletGoldsParams{
			UserId: params.OriginUserId,
			Amount: rewards.FromInvites.Metrics[4].Value,
		})
		if err != nil {
			logger.Error(err.Error())
			return "", err
		}
	}

	return "", nil
}
