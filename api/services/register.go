package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/configs"
)

type RegisterService interface {
	RegisterUser(ctx context.Context, name string, username string, email string) error
}

type registerService struct {
	us UserService
	ss SessionService
}

func NewRegisterService(
	userService UserService,
	sessionService SessionService) RegisterService {
	return &registerService{
		us: userService,
		ss: sessionService,
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

	// TODO: Send email with magic link
	fmt.Println("Magic link:", magicLink)

	return nil
}
