package services

import (
	"context"
	"fmt"

	"github.com/hermes-mailer/models"
)

type EmailService interface {
	Send(ctx context.Context, email models.Email) error
}

type emailService struct {
	emailClient models.EmailSenderClient
}

func NewEmailService(client models.EmailSenderClient) EmailService {
	return &emailService{
		emailClient: client,
	}
}

func (e *emailService) Send(ctx context.Context, email models.Email) error {
	if err := e.emailClient.SendEmail(ctx, email); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	// TODO: Save email to database

	return nil
}
