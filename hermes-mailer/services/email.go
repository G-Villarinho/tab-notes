package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hermes-mailer/config"
	"github.com/hermes-mailer/models"
)

type EmailService interface {
	Enqueue(ctx context.Context, email models.Email) error
	SendEmail(ctx context.Context, email models.Email) error
}

type emailService struct {
	ec models.EmailSenderClient
	qs QueueService
}

func NewEmailService(
	client models.EmailSenderClient,
	queueService QueueService) EmailService {
	return &emailService{
		ec: client,
		qs: queueService,
	}
}

func (e *emailService) SendEmail(ctx context.Context, email models.Email) error {
	if err := e.ec.SendEmail(ctx, email); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}

func (e *emailService) Enqueue(ctx context.Context, email models.Email) error {
	bytes, err := json.Marshal(email)
	if err != nil {
		return fmt.Errorf("marshal email: %w", err)
	}

	if err := e.qs.Publish(ctx, config.Env.API.QueueName, bytes); err != nil {
		return fmt.Errorf("enqueue email: %w", err)
	}

	return nil
}
