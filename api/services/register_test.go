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
		emailNotification := new(mocks.EmailNotificationMock)

		s := NewRegisterService(userService, sessionService, emailNotification)

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
		emailNotification := new(mocks.EmailNotificationMock)

		s := NewRegisterService(userService, sessionService, emailNotification)

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

	t.Run("should return error if SendWelcomeEmail fails", func(t *testing.T) {
		userService := new(mocks.UserServiceMock)
		sessionService := new(mocks.SessionServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)

		s := NewRegisterService(userService, sessionService, emailNotification)

		user := &models.User{ID: "user-123", Name: "João", Email: "joao@example.com"}

		userService.
			On("CreateUser", ctx, "João", "joaozinha-30", "joao@example.com").
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, "user-123", "joao@example.com").
			Return("token123", nil)

		configs.Env.APIURL = "http://localhost:8080"
		expectedLink := "http://localhost:8080/magic-link/authenticate?token=token123"

		emailNotification.
			On("SendWelcomeEmail", ctx, "João", "joao@example.com", expectedLink).
			Return(errors.New("email fail"))

		err := s.RegisterUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.ErrorContains(t, err, "email fail")
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
		emailNotification.AssertExpectations(t)
	})

	t.Run("should create user, session and send welcome email successfully", func(t *testing.T) {
		userService := new(mocks.UserServiceMock)
		sessionService := new(mocks.SessionServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)

		s := NewRegisterService(userService, sessionService, emailNotification)

		user := &models.User{ID: "user-123", Name: "João", Email: "joao@example.com"}

		userService.
			On("CreateUser", ctx, "João", "joaozinha-30", "joao@example.com").
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, "user-123", "joao@example.com").
			Return("mock-token-abc", nil)

		configs.Env.APIURL = "http://localhost:8080"
		expectedLink := fmt.Sprintf("%s/magic-link/authenticate?token=%s", configs.Env.APIURL, "mock-token-abc")

		emailNotification.
			On("SendWelcomeEmail", ctx, "João", "joao@example.com", expectedLink).
			Return(nil)

		err := s.RegisterUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.NoError(t, err)
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
		emailNotification.AssertExpectations(t)
	})
}
