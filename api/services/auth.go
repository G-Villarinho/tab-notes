package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/models"
)

type AuthService interface {
	SendAuthenticationLink(ctx context.Context, email string) error
	AuthenticateFromLink(ctx context.Context, token string) (*models.AuthResponse, error)
	Logout(ctx context.Context, sessionId string) error
}

type authService struct {
	ss SessionService
	us UserService
	qs QueueService
}

func NewAuthService(
	sessionService SessionService,
	userService UserService,
	queueService QueueService) AuthService {
	return &authService{
		ss: sessionService,
		us: userService,
		qs: queueService,
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

	body := fmt.Sprintf("Olá %s! 👋\n\nClique no link abaixo para acessar sua conta:\n\n%s\n\nEste link é válido por tempo limitado.", user.Name, magicLink)

	if err := a.qs.SendEmail(ctx, email, "Acesse sua conta no Tab Notes", body); err != nil {
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
