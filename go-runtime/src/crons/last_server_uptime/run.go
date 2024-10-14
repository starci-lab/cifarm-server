package crons_last_server_uptime

import (
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
	//delay 10s to start write
	time.Sleep(10 * time.Second)

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
			func() {
				Process(ctx, logger, db, nk)
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
