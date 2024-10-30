package hooks

import (
	collections_buildings "cifarm-server/src/collections/buildings"
	collections_common "cifarm-server/src/collections/common"
	collections_placed_items "cifarm-server/src/collections/placed_items"
	collections_player "cifarm-server/src/collections/player"
	collections_system "cifarm-server/src/collections/system"
	collections_tiles "cifarm-server/src/collections/tiles"
	"cifarm-server/src/utils"
	"cifarm-server/src/wallets"
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type HandleRefererParams struct {
	//your metadata
	Metadata       collections_player.Metadata `json:"metadata"`
	ReferrerUserId string                      `json:"referrerUserId"`
}

func HandleReferer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, userId string, params HandleRefererParams) error {
	object, err := collections_system.ReadRewards(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "rewards not found"
		logger.Error(errMsg)
		return err
	}
	rewards, err := collections_common.ToValue[collections_system.Rewards](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	object, err = collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: params.ReferrerUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		debugMsg := "player stats not found"
		//if not found, mean wrong code, stop the refer
		logger.Debug(debugMsg)
		return nil
	}
	playerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	//check
	if utils.ContainsInt(playerStats.Invites, params.Metadata.TelegramData.UserId) {
		debugMsg := "already invited"
		logger.Debug(debugMsg)
		return nil
	}

	//bonus for being referred
	err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
		GoldAmount: rewards.Referred,
		Metadata: map[string]interface{}{
			"name": "Referred",
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	playerStats.Invites = append(playerStats.Invites, params.Metadata.TelegramData.UserId)

	err = collections_player.WritePlayerStats(ctx, logger, db, nk, collections_player.WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      params.ReferrerUserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if len(playerStats.Invites) == rewards.FromInvites.Metrics[1].Key {
		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     params.ReferrerUserId,
			GoldAmount: rewards.FromInvites.Metrics[1].Value,
			Metadata: map[string]interface{}{
				"name": "Refer",
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[2].Key {
		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     params.ReferrerUserId,
			GoldAmount: rewards.FromInvites.Metrics[2].Value,
			Metadata: map[string]interface{}{
				"name": "Refer",
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[3].Key {
		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     params.ReferrerUserId,
			GoldAmount: rewards.FromInvites.Metrics[3].Value,
			Metadata: map[string]interface{}{
				"name": "Refer",
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	} else if len(playerStats.Invites) == rewards.FromInvites.Metrics[4].Key {
		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			UserId:     params.ReferrerUserId,
			GoldAmount: rewards.FromInvites.Metrics[4].Value,
			Metadata: map[string]interface{}{
				"name": "Refer",
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	return nil
}

type Claims struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func AfterAuthenticate(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	out *api.Session,
	in *api.AuthenticateCustomRequest,
) error {
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		errMsg := "user ID not found"
		logger.Error(errMsg)
		return errors.New(errMsg)
	}

	chain := in.Account.Vars["chainKey"]
	accountAddress := in.Account.Vars["accountAddress"]
	network := in.Account.Vars["network"]
	telegramUserId := in.Account.Vars["telegramUserId"]
	_telegramUserId, err := strconv.Atoi(telegramUserId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	object, err := collections_player.ReadMetadata(ctx, logger, db, nk, collections_player.ReadMetadataParams{
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	//when user auth, reset visit state to home
	err = collections_player.WriteVisitState(ctx, logger, db, nk, collections_player.WriteVisitStateParams{
		VisitState: collections_player.VisitState{
			UserId: "",
		},
		UserId: userId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	if object == nil {
		//get
		object, err := collections_system.ReadStarterConfigure(ctx, logger, db, nk)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		starterConfigure, err := collections_common.ToValue[collections_system.StarterConfigure](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		//first time login
		metadata := collections_player.Metadata{
			ChainKey:       chain,
			AccountAddress: accountAddress,
			Network:        network,
			TelegramData: collections_player.TelegramData{
				UserId: _telegramUserId,
			},
		}
		err = collections_player.WriteMetadata(ctx, logger, db, nk,
			collections_player.WriteMetadataParams{
				Metadata: metadata,
				UserId:   userId,
			})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = collections_player.WritePlayerStats(ctx, logger, db, nk, collections_player.WritePlayerStatsParams{
			UserId: userId,
			PlayerStats: collections_player.PlayerStats{
				LevelInfo: collections_player.LevelInfo{
					Level:           1,
					Experiences:     0,
					ExperienceQuota: 50,
				},
				TutorialInfo: collections_player.TutorialInfo{
					TutorialIndex: 0,
					StepIndex:     0,
				},
				EnergyInfo: collections_player.EnergyInfo{
					CurrentEnergy:     50,
					MaxEnergy:         50,
					EnergyQuota:       1,
					RecoveryTimeCount: 0,
				},
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = collections_player.WriteRewardTracker(ctx, logger, db, nk, collections_player.WriteRewardTrackerParams{
			UserId: userId,
			RewardTracker: collections_player.RewardTracker{
				DailyRewardsInfo: collections_player.DailyRewardsInfo{
					Streak:         0,
					LastClaimTime:  0,
					NumberOfClaims: 0,
				},
				SpinInfo: collections_player.SpinInfo{
					LastSpinTime: 0,
					SpinCount:    0,
				},
			}})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = collections_player.WriteFollowings(ctx, logger, db, nk, collections_player.WriteFollowingsParams{})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		positions := []collections_placed_items.Position{
			{X: 0, Y: -1},
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: -1},
			{X: 1, Y: 0},
			{X: 1, Y: 1},
		}

		var placedItems []collections_placed_items.PlacedItem
		for _, position := range positions {
			placedItems = append(placedItems, collections_placed_items.PlacedItem{
				ReferenceKey: collections_tiles.KEY_STARTER_TILE,
				Position:     position,
				Type:         collections_placed_items.TYPE_TILE,
			})
		}
		placedItems = append(placedItems, collections_placed_items.PlacedItem{
			ReferenceKey: collections_buildings.KEY_HOME,
			Position: collections_placed_items.Position{
				X: 4,
				Y: 0,
			},
			Type: collections_placed_items.TYPE_BUILDING,
		})

		err = collections_placed_items.WriteMany(ctx, logger, db, nk, collections_placed_items.WriteManyParams{
			PlacedItems: placedItems,
			UserId:      userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		err = wallets.UpdateWallet(ctx, logger, db, nk, wallets.UpdateWalletParams{
			GoldAmount: starterConfigure.GoldAmount,
			Metadata: map[string]interface{}{
				"name": "Initial",
			},
			UserId: userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		object, err = collections_system.ReadUsers(ctx, logger, db, nk)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		users, err := collections_common.ToValue[collections_system.Users](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		users.UserIds = append(users.UserIds, userId)
		err = collections_system.WriteUsers(ctx, logger, db, nk, collections_system.WriteUsersParams{
			Users: *users,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		referrerUserId := in.Account.Vars["referrerUserId"]
		//you have been referred by someone, we only check newly telegram
		logger.Debug("referrerUserId: %s", referrerUserId)
		if referrerUserId != "" {
			err = HandleReferer(ctx, logger, db, nk, userId, HandleRefererParams{
				ReferrerUserId: referrerUserId,
				Metadata:       metadata,
			})
			if err != nil {
				logger.Error(err.Error())
				return err
			}
		}
	} else {
		//update tele metadata if neccessary
		//check metadata
		metadataObject, err := collections_player.ReadMetadata(ctx, logger, db, nk, collections_player.ReadMetadataParams{
			UserId: userId,
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		if metadataObject != nil {
			//has metadata before
			//check if telegramUserId is the same
			metadata, err := collections_common.ToValue[collections_player.Metadata](ctx, logger, db, nk, metadataObject)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			if metadata.TelegramData.UserId != _telegramUserId {
				//telegramUserId is different, do update
				metadata.TelegramData.UserId = _telegramUserId
				err = collections_player.WriteMetadata(ctx, logger, db, nk, collections_player.WriteMetadataParams{
					Metadata: *metadata,
					UserId:   userId,
				})
				if err != nil {
					logger.Error(err.Error())
					return err
				}
			}
		}
	}

	return nil
}
