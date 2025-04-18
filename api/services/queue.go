package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/clients"
	"github.com/g-villarinho/tab-notes-api/models"
)

type QueueService interface {
	SendEmail(ctx context.Context, to, subject, body string) error
}

type queueService struct {
	queue clients.QueueClient
}

func NewQueueService(queue clients.QueueClient) QueueService {
	return &queueService{
		queue: queue,
	}
}

func (q *queueService) SendEmail(ctx context.Context, to, subject, body string) error {
	payload := models.EmailQueue{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	err := q.queue.Publish(ctx, "email_queue", payload)
	if err != nil {
		return fmt.Errorf("publish email to queue: %w", err)
	}

	return nil
}
