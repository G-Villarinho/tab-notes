package clients

import (
	"context"
	"fmt"
	"net/smtp"

	"github.com/hermes-mailer/config"
	"github.com/hermes-mailer/models"
)

type SMTPEmailSenderClient struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewSMTPEmailSenderClient() models.EmailSenderClient {
	return &SMTPEmailSenderClient{
		Host:     config.Env.SMTP.Host,
		Port:     config.Env.SMTP.Port,
		Username: config.Env.SMTP.Username,
		Password: config.Env.SMTP.Password,
	}
}

func (s *SMTPEmailSenderClient) SendEmail(ctx context.Context, email models.Email) error {
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)

	boundary := "hermes-boundary"

	msg := []byte("To: " + email.To + "\r\n" +
		"From: " + s.Username + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: multipart/alternative; boundary=" + boundary + "\r\n" +
		"\r\n" +
		"--" + boundary + "\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		email.BodyText + "\r\n" +
		"--" + boundary + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		email.BodyHTML + "\r\n" +
		"--" + boundary + "--")

	if err := smtp.SendMail(addr, auth, s.Username, []string{email.To}, msg); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}
