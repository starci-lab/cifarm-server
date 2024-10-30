package crons_energy_gain

import (
	collections_common "cifarm-server/src/collections/common"
	collections_player "cifarm-server/src/collections/player"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HandleEnergyGainParams struct {
	UserId        string `json:"userId"`
	TimeInSeconds int64  `json:"timeInSeconds"`
}

func HandleEnergyGain(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params HandleEnergyGainParams,

) error {
	object, err := collections_player.ReadPlayerStats(ctx, logger, db, nk, collections_player.ReadPlayerStatsParams{
		UserId: params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object == nil {
		errMsg := "player stats not found"
		logger.Error(errMsg)
		return err
	}
	playerStats, err := collections_common.ToValue[collections_player.PlayerStats](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	//5 min = 300 sec
	recoveryTime := int64(300)
	// max energy, not restore energy
	if playerStats.EnergyInfo.CurrentEnergy >= playerStats.EnergyInfo.MaxEnergy {
		return nil
	}

	playerStats.EnergyInfo.RecoveryTimeCount += params.TimeInSeconds
	for {
		if playerStats.EnergyInfo.RecoveryTimeCount >= recoveryTime {
			playerStats.EnergyInfo.RecoveryTimeCount -= recoveryTime
			playerStats.EnergyInfo.CurrentEnergy++
			if playerStats.EnergyInfo.CurrentEnergy >= playerStats.EnergyInfo.MaxEnergy {
				playerStats.EnergyInfo.CurrentEnergy = playerStats.EnergyInfo.MaxEnergy
				playerStats.EnergyInfo.RecoveryTimeCount = 0
				break
			}
		} else {
			break
		}
	}
	err = collections_player.WritePlayerStats(ctx, logger, db, nk, collections_player.WritePlayerStatsParams{
		PlayerStats: *playerStats,
		UserId:      params.UserId,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func Process(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	timeSinceLastUptime int64,
) error {
	object, err := collections_system.ReadUsers(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	users, err := collections_common.ToValue[collections_system.Users](ctx, logger, db, nk, object)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	var speedUpTime int64
	object, err = collections_system.ReadSpeedUp(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if object != nil {
		speedUp, err := collections_common.ToValue[collections_system.SpeedUp](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		speedUpTime = speedUp.SeedGrowthTime
	}
	if speedUpTime > 0 {
		err := collections_system.WriteSpeedUp(ctx, logger, db, nk, collections_system.WriteSpeedUpParams{
			SpeedUp: collections_system.SpeedUp{
				EnergyGain: 0,
			},
		})
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	for _, userId := range users.UserIds {
		go HandleEnergyGain(ctx, logger, db, nk, HandleEnergyGainParams{
			UserId:        userId,
			TimeInSeconds: 1 + timeSinceLastUptime + speedUpTime,
		})
	}

	return err
}
