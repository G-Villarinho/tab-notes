package models

import "context"

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailSenderClient interface {
	SendEmail(ctx context.Context, email Email) error
}
