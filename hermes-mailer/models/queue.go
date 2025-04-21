package models

import (
	"context"

	"github.com/streadway/amqp"
)

type QueueConsumer interface {
	Consume(queueName string) (<-chan amqp.Delivery, error)
	Publish(ctx context.Context, queueName string, body []byte) error
	Close()
}
