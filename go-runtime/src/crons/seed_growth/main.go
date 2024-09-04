package seed_growth

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/heroiclabs/nakama-common/runtime"
)

type RunSeedGrowthCronParams struct {
	UserId string `json:"userId"`
}

func RunSeedGrowthCron(
	ctx context.Context,
	logger runtime.Logger,
	db *sql.DB,
	nk runtime.NakamaModule,
	params RunSeedGrowthCronParams,
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
				logger.Info("%s: seeds growing...", params.UserId)
				HandleSeedGrowth(ctx, logger, db, nk, HandleSeedGrowthParams{
					UserId: params.UserId,
				})
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
