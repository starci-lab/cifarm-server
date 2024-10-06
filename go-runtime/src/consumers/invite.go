package consumers

import (
	services_kafka "cifarm-server/src/services/kafka-client"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	GROUP_ID = "invite-group"
	TOPIC    = "invite"
)

func PoolInviteConsumer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	consumer, err := services_kafka.NewConsumer(ctx, logger, db, nk, services_kafka.NewConsumerParams{
		GroupId: GROUP_ID,
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = consumer.Subscribe(TOPIC, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	run := true
	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			{
				logger.Error("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
			}
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
		consumer.Close()
	}
	return nil
}
