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

func TestSendAuthenticationLink(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if user not found", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		userService.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, errors.New("not found"))

		err := auth.SendAuthenticationLink(ctx, "joao@example.com")

		assert.ErrorContains(t, err, "not found")
		userService.AssertExpectations(t)
	})

	t.Run("should return error if CreateSession fails", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		user := &models.User{ID: "user-1", Name: "João", Email: "joao@example.com"}

		userService.
			On("GetUserByEmail", ctx, user.Email).
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, user.ID, user.Email).
			Return("", errors.New("session fail"))

		err := auth.SendAuthenticationLink(ctx, user.Email)

		assert.ErrorContains(t, err, "session fail")
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
	})

	t.Run("should send email with magic link", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		user := &models.User{ID: "user-1", Name: "João", Email: "joao@example.com"}

		userService.
			On("GetUserByEmail", ctx, user.Email).
			Return(user, nil)

		sessionService.
			On("CreateSession", ctx, user.ID, user.Email).
			Return("token-xyz", nil)

		configs.Env.APIURL = "http://localhost:8080"

		expectedLink := fmt.Sprintf("%s/magic-link/authenticate?token=%s", configs.Env.APIURL, "token-xyz")

		emailNotification.
			On("SendMagicLink", ctx, user.Email, user.Name, expectedLink).
			Return(nil)

		err := auth.SendAuthenticationLink(ctx, user.Email)

		assert.NoError(t, err)
		userService.AssertExpectations(t)
		sessionService.AssertExpectations(t)
		emailNotification.AssertExpectations(t)
	})
}
func TestAuthenticateFromLink(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if session is invalid", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		sessionService.
			On("ValidSession", ctx, "invalid-token").
			Return("", errors.New("invalid or expired"))

		resp, err := auth.AuthenticateFromLink(ctx, "invalid-token")

		assert.Nil(t, resp)
		assert.ErrorContains(t, err, "invalid or expired")
		sessionService.AssertExpectations(t)
	})

	t.Run("should return auth response if session is valid", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		sessionService.
			On("ValidSession", ctx, "valid-token").
			Return("new-auth-token", nil)

		resp, err := auth.AuthenticateFromLink(ctx, "valid-token")

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "new-auth-token", resp.Token)
		sessionService.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if session revocation fails", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		sessionService.
			On("RevokeSession", ctx, "sess-123").
			Return(errors.New("fail to revoke"))

		err := auth.Logout(ctx, "sess-123")

		assert.ErrorContains(t, err, "fail to revoke")
		sessionService.AssertExpectations(t)
	})

	t.Run("should logout successfully", func(t *testing.T) {
		sessionService := new(mocks.SessionServiceMock)
		userService := new(mocks.UserServiceMock)
		emailNotification := new(mocks.EmailNotificationMock)
		auth := NewAuthService(sessionService, userService, emailNotification)

		sessionService.
			On("RevokeSession", ctx, "sess-456").
			Return(nil)

		err := auth.Logout(ctx, "sess-456")

		assert.NoError(t, err)
		sessionService.AssertExpectations(t)
	})
}
