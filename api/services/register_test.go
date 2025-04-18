package services

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if CreateUser fails", func(t *testing.T) {
		userService := new(mocks.UserServiceMock)
		sessionService := new(mocks.SessionServiceMock)
		s := NewRegisterService(userService, sessionService)

		userService.
			On("CreateUser", ctx, "João", "joaozinha-30", "joao@example.com").
			Return(nil, errors.New("db error"))

		err := s.RegisterUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.ErrorContains(t, err, "db error")
		userService.AssertExpectations(t)
	})

	t.Run("should return error if CreateSession fails", func(t *testing.T) {
		userService := new(mocks.UserServiceMock)
		sessionService := new(mocks.SessionServiceMock)
		s := NewRegisterService(userService, sessionService)

		user := &models.User{ID: "user-123", Name: "João", Email: "joao@example.com"}

		userService.
			On("CreateUser", ctx, "João", "joaozinha-30", "joao@example.com").
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, "user-123", "joao@example.com").
			Return("", errors.New("session error"))

		err := s.RegisterUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.ErrorContains(t, err, "session error")
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
	})

	t.Run("should create user and session successfully", func(t *testing.T) {
		userService := new(mocks.UserServiceMock)
		sessionService := new(mocks.SessionServiceMock)
		s := NewRegisterService(userService, sessionService)

		user := &models.User{ID: "user-123", Name: "João", Email: "joao@example.com"}

		userService.
			On("CreateUser", ctx, "João", "joaozinha-30", "joao@example.com").
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, "user-123", "joao@example.com").
			Return("mock-token-abc", nil)

		configs.Env.APIURL = "http://localhost:8080"

		expectedLink := fmt.Sprintf("%s/magic-link/authenticate?token=%s", configs.Env.APIURL, "mock-token-abc")

		t.Log("Esperado:", expectedLink)

		err := s.RegisterUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.NoError(t, err)
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
	})
}
