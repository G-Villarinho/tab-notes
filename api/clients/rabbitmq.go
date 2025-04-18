package clients

import (
	"context"
	"encoding/json"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/streadway/amqp"
)

type QueueClient interface {
	Publish(ctx context.Context, queue string, payload any) error
	Close()
}

type rabbitMQPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQPublisher() (QueueClient, error) {
	conn, err := amqp.Dial(configs.Env.AMQPURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &rabbitMQPublisher{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *rabbitMQPublisher) Publish(ctx context.Context, queue string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = r.ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return r.ch.Publish(
			"",
			queue,
			false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		)
	}
}

func (r *rabbitMQPublisher) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}
