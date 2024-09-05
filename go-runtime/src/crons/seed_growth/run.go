package crons_seed_growth

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/heroiclabs/nakama-common/runtime"
)

func RunSeedGrowthCron(
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
		gocron.DurationJob(
			time.Second,
		),
		gocron.NewTask(
			func() {
				logger.Info("seeds growing...")
				Process(ctx, logger, db, nk)
			},
		),
	)

	logger.Info("seed growth job: %s", job.ID())

	if err != nil {
		logger.Error(err.Error())
		return err
	}
	scheduler.Start()
	return nil
}
