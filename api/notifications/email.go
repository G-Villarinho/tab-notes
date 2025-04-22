package notifications

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/g-villarinho/tab-notes-api/clients"
	"github.com/g-villarinho/tab-notes-api/models"
)

type EmailNotification interface {
	SendMagicLink(ctx context.Context, name string, email string, magicLink string) error
	SendWelcomeEmail(ctx context.Context, name, email, magicLink string) error
}

type emailNotification struct {
	ec   clients.EmailClient
	path string
}

func NewEmailNotification(ec clients.EmailClient) EmailNotification {
	return &emailNotification{
		ec:   ec,
		path: "notifications/templates",
	}
}
func (e *emailNotification) SendMagicLink(ctx context.Context, name string, email string, magicLink string) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/access-link-email.html", e.path))
	if err != nil {
		log.Fatalf("parse template: %v", err)
	}

	var htmlBuffer bytes.Buffer
	data := models.MagicLinkEmailData{
		MagicLink: magicLink,
		Name:      name,
		Year:      time.Now().Year(),
	}

	if err := tmpl.Execute(&htmlBuffer, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	emailData := &models.Email{
		To:       email,
		Subject:  "Login to Tab Notes",
		BodyText: fmt.Sprintf("Hello %s,\n\nClick the link below to login to Tab Notes:\n%s\n\nBest regards,\nTab Notes Team", name, magicLink),
		BodyHTML: htmlBuffer.String(),
	}

	if err := e.ec.SendEmail(ctx, emailData); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}

func (e *emailNotification) SendWelcomeEmail(ctx context.Context, name, email, magicLink string) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/welcome-access-email.html", e.path))
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	var htmlBuffer bytes.Buffer
	data := models.MagicLinkEmailData{
		MagicLink: magicLink,
		Name:      name,
		Year:      time.Now().Year(),
	}

	if err := tmpl.Execute(&htmlBuffer, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	emailData := &models.Email{
		To:       email,
		Subject:  "Bem-vindo ao Tab Notes!",
		BodyText: fmt.Sprintf("Olá %s!\n\nClique no link abaixo para acessar sua conta:\n%s\n\nBem-vindo(a) ao Tab Notes!", name, magicLink),
		BodyHTML: htmlBuffer.String(),
	}

	return e.ec.SendEmail(ctx, emailData)
}
