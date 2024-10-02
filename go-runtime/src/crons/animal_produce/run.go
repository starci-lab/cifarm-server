package crons_animal_produce

import (
	collections_common "cifarm-server/src/collections/common"
	collections_system "cifarm-server/src/collections/system"
	"context"
	"database/sql"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/heroiclabs/nakama-common/runtime"
)

func Run(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	object, err := collections_system.ReadLastServerUptime(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	var timeSinceLastUptime int64
	if object != nil {
		lastServerUptime, err := collections_common.ToValue[collections_system.LastServerUptime](ctx, logger, db, nk, object)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		timeSinceLastUptime = time.Now().Unix() - lastServerUptime.TimeInSeconds

		logger.Info("time since last uptime: %vs", timeSinceLastUptime)
	}

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	job, err := scheduler.NewJob(
		gocron.DurationJob(
			time.Second,
		),
		gocron.NewTask(
			func() error {
				go Process(ctx, logger, db, nk, timeSinceLastUptime)
				if timeSinceLastUptime > 0 {
					timeSinceLastUptime = 0
				}
				return nil
			},
		),
	)
	logger.Info(`Job started: %s"`, job.ID())
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	scheduler.Start()
	return nil
}
