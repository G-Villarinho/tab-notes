package services

import (
	"context"

	"github.com/hermes-mailer/models"
)

type QueueService interface {
	Publish(ctx context.Context, queueName string, body []byte) error
	Consume(ctx context.Context, queueName string) (<-chan []byte, error)
}

type queueService struct {
	queueClient models.QueueConsumer
}

func NewQueueService(
	client models.QueueConsumer,
	queueName string) QueueService {
	return &queueService{
		queueClient: client,
	}
}

func (q *queueService) Publish(ctx context.Context, queueName string, body []byte) error {
	err := q.queueClient.Publish(ctx, queueName, body)
	if err != nil {
		return err
	}

	return nil
}

func (q *queueService) Consume(ctx context.Context, queueName string) (<-chan []byte, error) {
	deliveries, err := q.queueClient.Consume(queueName)
	if err != nil {
		return nil, err
	}

	byteChan := make(chan []byte)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(byteChan)
				return
			case msg, ok := <-deliveries:
				if !ok {
					close(byteChan)
					return
				}
				byteChan <- msg.Body
			}
		}
	}()

	return byteChan, nil
}
