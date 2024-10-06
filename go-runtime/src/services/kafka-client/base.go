package services_kafka

import (
	"cifarm-server/src/config"
	"context"
	"database/sql"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/heroiclabs/nakama-common/runtime"
)

func NewProducer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (*kafka.Producer, error) {
	host, err := config.Kafka1Host(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	port, err := config.Kafka1Port(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"client.id":         "cifarm-server",
		"acks":              "all",
	})

	if err != nil {
		return nil, err
	}
	return producer, nil
}

type NewConsumerParams struct {
	GroupId string `json:"groupId"`
}

func NewConsumer(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params NewConsumerParams) (*kafka.Consumer, error) {
	host, err := config.Kafka1Host(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	port, err := config.Kafka1Port(ctx, logger, db, nk)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"client.id":         "cifarm-server",
		"group.id":          params.GroupId,
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		return nil, err
	}
	return consumer, nil
}
