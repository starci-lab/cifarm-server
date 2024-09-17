package crons_deliver

import (
	"context"
	"database/sql"

	"github.com/go-co-op/gocron/v2"
	"github.com/heroiclabs/nakama-common/runtime"
)

func Run(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
) error {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	job, err := scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(
			gocron.NewAtTime(0, 0, 0),
		)),
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
