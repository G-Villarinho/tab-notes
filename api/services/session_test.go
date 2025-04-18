package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func FakeSession(userID string, token string) *models.Session {
	now := time.Now().UTC()
	return &models.Session{
		ID:        "sess-123",
		UserID:    userID,
		Token:     token,
		CreatedAt: now,
		ExpiresAt: now.Add(15 * time.Minute),
	}
}

func TestCreateSession(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if token generation fails", func(t *testing.T) {
		tokenService := new(mocks.TokenServiceMock)
		sessionRepo := new(mocks.SessionRepositoryMock)
		s := NewSessionService(tokenService, sessionRepo)

		tokenService.
			On("GenerateMagicLinkToken", ctx, "joao@example.com", pkgs.MockAnyTime(), pkgs.MockAnyTime()).
			Return("", errors.New("fail to sign token"))

		token, err := s.CreateSession(ctx, "user-1", "joao@example.com")

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "fail to sign token")
		tokenService.AssertExpectations(t)
	})

	t.Run("should return error if session creation fails", func(t *testing.T) {
		tokenService := new(mocks.TokenServiceMock)
		sessionRepo := new(mocks.SessionRepositoryMock)
		s := NewSessionService(tokenService, sessionRepo)

		tokenService.
			On("GenerateMagicLinkToken", ctx, "joao@example.com", pkgs.MockAnyTime(), pkgs.MockAnyTime()).
			Return("fake-jwt-token", nil)

		sessionRepo.
			On("CreateSession", ctx, pkgs.MockSessionWithToken("fake-jwt-token")).
			Return(errors.New("db error"))

		token, err := s.CreateSession(ctx, "user-1", "joao@example.com")

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "db error")
		tokenService.AssertExpectations(t)
		sessionRepo.AssertExpectations(t)
	})

	t.Run("should return token on success", func(t *testing.T) {
		tokenService := new(mocks.TokenServiceMock)
		sessionRepo := new(mocks.SessionRepositoryMock)
		s := NewSessionService(tokenService, sessionRepo)

		tokenService.
			On("GenerateMagicLinkToken", ctx, "joao@example.com", pkgs.MockAnyTime(), pkgs.MockAnyTime()).
			Return("valid-token", nil)

		sessionRepo.
			On("CreateSession", ctx, pkgs.MockSessionWithToken("valid-token")).
			Return(nil)

		token, err := s.CreateSession(ctx, "user-1", "joao@example.com")

		assert.NoError(t, err)
		assert.Equal(t, "valid-token", token)
		tokenService.AssertExpectations(t)
		sessionRepo.AssertExpectations(t)
	})
}

func TestValidSession(t *testing.T) {
	ctx := context.Background()
	now := time.Now().UTC()

	t.Run("should return error if repo fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("GetSessionByToken", ctx, "token-123").
			Return(nil, errors.New("repo fail"))

		token, err := s.ValidSession(ctx, "token-123")

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "repo fail")
		sr.AssertExpectations(t)
	})

	t.Run("should return ErrSessionNotFound", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("GetSessionByToken", ctx, "token-abc").
			Return(nil, nil)

		token, err := s.ValidSession(ctx, "token-abc")

		assert.Empty(t, token)
		assert.ErrorIs(t, err, models.ErrSessionNotFound)
		sr.AssertExpectations(t)
	})

	t.Run("should return ErrSessionExpired and delete session", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "token-expired")
		session.ExpiresAt = now.Add(-1 * time.Minute)

		sr.On("GetSessionByToken", ctx, "token-expired").
			Return(session, nil)

		sr.On("DeleteSession", ctx, session.ID).
			Return(nil)

		token, err := s.ValidSession(ctx, "token-expired")

		assert.Empty(t, token)
		assert.ErrorIs(t, err, models.ErrSessionExpired)
		sr.AssertExpectations(t)
	})

	t.Run("should return error if token generation fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "magic-token")

		sr.On("GetSessionByToken", ctx, "magic-token").
			Return(session, nil)

		ts.On("GenerateAuthToken", ctx, session.UserID, session.ID, session.CreatedAt, mock.Anything).
			Return("", errors.New("sign error"))

		token, err := s.ValidSession(ctx, "magic-token")

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "sign error")
		sr.AssertExpectations(t)
		ts.AssertExpectations(t)
	})

	t.Run("should return error if update fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "magic-token")

		sr.On("GetSessionByToken", ctx, "magic-token").
			Return(session, nil)

		ts.On("GenerateAuthToken", ctx, session.UserID, session.ID, session.CreatedAt, mock.Anything).
			Return("new-token", nil)

		sr.On("UpdateSession", ctx, mock.MatchedBy(func(s *models.Session) bool {
			return s.Token == "new-token"
		})).Return(errors.New("update error"))

		token, err := s.ValidSession(ctx, "magic-token")

		assert.Empty(t, token)
		assert.ErrorContains(t, err, "update error")
	})

	t.Run("should return auth token on success", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "magic-token")

		sr.On("GetSessionByToken", ctx, "magic-token").
			Return(session, nil)

		ts.On("GenerateAuthToken", ctx, session.UserID, session.ID, session.CreatedAt, mock.Anything).
			Return("new-auth-token", nil)

		sr.On("UpdateSession", ctx, mock.MatchedBy(func(s *models.Session) bool {
			return s.Token == "new-auth-token" && s.VerifiedAt.Valid
		})).Return(nil)

		token, err := s.ValidSession(ctx, "magic-token")

		assert.NoError(t, err)
		assert.Equal(t, "new-auth-token", token)
		sr.AssertExpectations(t)
		ts.AssertExpectations(t)
	})
}

func TestRevokeSession(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if GetSessionById fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("GetSessionById", ctx, "sess-1").
			Return(nil, errors.New("db error"))

		err := s.RevokeSession(ctx, "sess-1")

		assert.ErrorContains(t, err, "db error")
		sr.AssertExpectations(t)
	})

	t.Run("should return ErrSessionNotFound if session does not exist", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("GetSessionById", ctx, "sess-2").
			Return(nil, nil)

		err := s.RevokeSession(ctx, "sess-2")

		assert.ErrorIs(t, err, models.ErrSessionNotFound)
		sr.AssertExpectations(t)
	})

	t.Run("should return error if RevokeSession fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "token-123")

		sr.On("GetSessionById", ctx, session.ID).
			Return(session, nil)

		sr.On("RevokeSession", ctx, session.ID, mock.AnythingOfType("time.Time")).
			Return(errors.New("revoke error"))

		err := s.RevokeSession(ctx, session.ID)

		assert.ErrorContains(t, err, "revoke error")
		sr.AssertExpectations(t)
	})

	t.Run("should revoke session successfully", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		session := FakeSession("user-1", "token-123")

		sr.On("GetSessionById", ctx, session.ID).
			Return(session, nil)

		sr.On("RevokeSession", ctx, session.ID, mock.AnythingOfType("time.Time")).
			Return(nil)

		err := s.RevokeSession(ctx, session.ID)

		assert.NoError(t, err)
		sr.AssertExpectations(t)
	})
}

func TestIsSessionRevoked(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repo fails", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("IsSessionRevoked", ctx, "sess-1").
			Return(false, errors.New("repo error"))

		revoked, err := s.IsSessionRevoked(ctx, "sess-1")

		assert.False(t, revoked)
		assert.ErrorContains(t, err, "check if session sess-1 is revoked: repo error")
		sr.AssertExpectations(t)
	})

	t.Run("should return true if session is revoked", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("IsSessionRevoked", ctx, "sess-2").
			Return(true, nil)

		revoked, err := s.IsSessionRevoked(ctx, "sess-2")

		assert.NoError(t, err)
		assert.True(t, revoked)
		sr.AssertExpectations(t)
	})

	t.Run("should return false if session is not revoked", func(t *testing.T) {
		ts := new(mocks.TokenServiceMock)
		sr := new(mocks.SessionRepositoryMock)
		s := NewSessionService(ts, sr)

		sr.On("IsSessionRevoked", ctx, "sess-3").
			Return(false, nil)

		revoked, err := s.IsSessionRevoked(ctx, "sess-3")

		assert.NoError(t, err)
		assert.False(t, revoked)
		sr.AssertExpectations(t)
	})
}
