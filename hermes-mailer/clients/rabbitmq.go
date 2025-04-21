package clients

import (
	"context"

	"github.com/hermes-mailer/models"
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQClient(amqpURL string) (models.QueueConsumer, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQClient) Consume(queueName string) (<-chan amqp.Delivery, error) {
	_, err := r.ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	msgs, err := r.ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMQClient) Publish(ctx context.Context, queueName string, body []byte) error {
	_, err := r.ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return r.ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMQClient) Close() {
	_ = r.ch.Close()
	_ = r.conn.Close()
}
