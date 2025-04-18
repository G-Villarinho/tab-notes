package models

import "github.com/streadway/amqp"

type QueueConsumer interface {
	Consume(queueName string) (<-chan amqp.Delivery, error)
	Close()
}
