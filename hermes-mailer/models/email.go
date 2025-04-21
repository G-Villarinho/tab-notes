package models

import "context"

type Email struct {
	To       string `json:"to"`
	Subject  string `json:"subject"`
	BodyText string `json:"body_text"`
	BodyHTML string `json:"body_html"`
}

type EmailSenderClient interface {
	SendEmail(ctx context.Context, email Email) error
}
