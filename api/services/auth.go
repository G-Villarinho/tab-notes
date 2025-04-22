package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/notifications"
)

type AuthService interface {
	SendAuthenticationLink(ctx context.Context, email string) error
	AuthenticateFromLink(ctx context.Context, token string) (*models.AuthResponse, error)
	Logout(ctx context.Context, sessionId string) error
}

type authService struct {
	ss SessionService
	us UserService
	en notifications.EmailNotification
}

func NewAuthService(
	sessionService SessionService,
	userService UserService,
	emailNotification notifications.EmailNotification,
) AuthService {
	return &authService{
		ss: sessionService,
		us: userService,
		en: emailNotification,
	}
}

func (a *authService) SendAuthenticationLink(ctx context.Context, email string) error {
	user, err := a.us.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	tokenMagicLink, err := a.ss.CreateSession(ctx, user.ID, email)
	if err != nil {
		return err
	}

	magicLink := fmt.Sprintf("%s/magic-link/authenticate?token=%s", configs.Env.APIURL, tokenMagicLink)

	if err := a.en.SendMagicLink(ctx, user.Name, user.Email, magicLink); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	return nil
}

func (a *authService) AuthenticateFromLink(ctx context.Context, token string) (*models.AuthResponse, error) {
	authToken, err := a.ss.ValidSession(ctx, token)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: authToken,
	}, nil
}

func (a *authService) Logout(ctx context.Context, sessionId string) error {
	if err := a.ss.RevokeSession(ctx, sessionId); err != nil {
		return err
	}

	return nil
}
