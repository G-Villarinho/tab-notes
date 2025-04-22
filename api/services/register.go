package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/notifications"
)

type RegisterService interface {
	RegisterUser(ctx context.Context, name string, username string, email string) error
}

type registerService struct {
	us UserService
	ss SessionService
	en notifications.EmailNotification
}

func NewRegisterService(
	userService UserService,
	sessionService SessionService,
	emailNotification notifications.EmailNotification) RegisterService {
	return &registerService{
		us: userService,
		ss: sessionService,
		en: emailNotification,
	}
}

func (r *registerService) RegisterUser(ctx context.Context, name string, username string, email string) error {
	user, err := r.us.CreateUser(ctx, name, username, email)
	if err != nil {
		return err
	}

	tokenMagicLink, err := r.ss.CreateSession(ctx, user.ID, email)
	if err != nil {
		return err
	}

	magicLink := fmt.Sprintf("%s/magic-link/authenticate?token=%s", configs.Env.APIURL, tokenMagicLink)

	if err := r.en.SendWelcomeEmail(ctx, name, email, magicLink); err != nil {
		return fmt.Errorf("send welcome email: %w", err)
	}

	return nil
}
