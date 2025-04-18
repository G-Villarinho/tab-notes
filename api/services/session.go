package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type SessionService interface {
	CreateSession(ctx context.Context, userID, email string) (string, error)
	ValidSession(ctx context.Context, token string) (string, error)
	RevokeSession(ctx context.Context, sessionId string) error
	IsSessionRevoked(ctx context.Context, sessionId string) (bool, error)
	GetUserSessions(ctx context.Context, userID string) ([]*models.SessionResponse, error)
	RevokeUserSession(ctx context.Context, userID string, sessionID string) error
	RevokeAllUserSessions(ctx context.Context, userID string, currentSessionID string, revokeCurrent bool) error
}

type sessionService struct {
	ts TokenService
	sr repositories.SessionRepository
}

func NewSessionService(tokenService TokenService, sessionRepository repositories.SessionRepository) SessionService {
	return &sessionService{
		ts: tokenService,
		sr: sessionRepository,
	}
}

func (s *sessionService) CreateSession(ctx context.Context, userID string, email string) (string, error) {
	now := time.Now().UTC()
	session := models.Session{
		UserID:    userID,
		CreatedAt: now,
		ExpiresAt: now.Add(time.Minute * 15),
	}

	tokenMagicLink, err := s.ts.GenerateMagicLinkToken(ctx, email, session.CreatedAt, session.ExpiresAt)
	if err != nil {
		return "", err
	}

	session.Token = tokenMagicLink

	if err := s.sr.CreateSession(ctx, &session); err != nil {
		return "", err
	}

	return session.Token, nil
}

func (s *sessionService) ValidSession(ctx context.Context, token string) (string, error) {
	now := time.Now().UTC()

	session, err := s.sr.GetSessionByToken(ctx, token)
	if err != nil {
		return "", err
	}

	if session == nil {
		return "", models.ErrSessionNotFound
	}

	if session.ExpiresAt.Before(now) {
		if err := s.sr.DeleteSession(ctx, session.ID); err != nil {
			return "", fmt.Errorf("delete session %s: %w", session.ID, err)
		}

		return "", models.ErrSessionExpired
	}

	expiresAt := now.Add(time.Hour * 24 * 7)

	authToken, err := s.ts.GenerateAuthToken(ctx, session.UserID, session.ID, session.CreatedAt, expiresAt)
	if err != nil {
		return "", err
	}

	session.Token = authToken
	session.VerifiedAt = sql.NullTime{Time: now, Valid: true}
	session.ExpiresAt = expiresAt

	if err := s.sr.UpdateSession(ctx, session); err != nil {
		return "", fmt.Errorf("update session %s: %w", session.ID, err)
	}

	return session.Token, nil
}

func (s *sessionService) RevokeSession(ctx context.Context, sessionId string) error {
	session, err := s.sr.GetSessionById(ctx, sessionId)
	if err != nil {
		return err
	}

	if session == nil {
		return models.ErrSessionNotFound
	}

	session.RevokedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if err := s.sr.RevokeSession(ctx, session.ID, session.RevokedAt.Time); err != nil {
		return fmt.Errorf("revoke session %s: %w", session.ID, err)
	}

	return nil
}

func (s *sessionService) IsSessionRevoked(ctx context.Context, sessionId string) (bool, error) {
	revoked, err := s.sr.IsSessionRevoked(ctx, sessionId)
	if err != nil {
		return false, fmt.Errorf("check if session %s is revoked: %w", sessionId, err)
	}

	return revoked, nil
}

func (s *sessionService) GetUserSessions(ctx context.Context, userID string) ([]*models.SessionResponse, error) {
	sessions, err := s.sr.GetSessionsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get sessions by user id: %w", err)
	}

	if len(sessions) == 0 {
		return []*models.SessionResponse{}, nil
	}

	var sessionResponses []*models.SessionResponse
	for _, session := range sessions {
		sessionResponse := &models.SessionResponse{
			ID:        session.ID,
			ExpiresAt: session.ExpiresAt,
			CreatedAt: session.CreatedAt,
		}

		if session.VerifiedAt.Valid {
			sessionResponse.VerifiedAt = &session.VerifiedAt.Time
		}

		if session.RevokedAt.Valid {
			sessionResponse.RevokedAt = &session.RevokedAt.Time
		}

		sessionResponses = append(sessionResponses, sessionResponse)
	}

	return sessionResponses, nil
}

func (s *sessionService) RevokeUserSession(ctx context.Context, userID string, sessionID string) error {
	session, err := s.sr.GetSessionById(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("get session by id %s: %w", sessionID, err)
	}

	if session == nil {
		return models.ErrSessionNotFound
	}

	if session.UserID != userID {
		return models.ErrSessionNotFound
	}

	if err := s.sr.RevokeSession(ctx, session.ID, time.Now().UTC()); err != nil {
		return fmt.Errorf("revoke session %s: %w", sessionID, err)
	}

	return nil
}

func (s *sessionService) RevokeAllUserSessions(ctx context.Context, userID string, currentSessionID string, revokeCurrent bool) error {
	now := time.Now().UTC()

	if revokeCurrent {
		if err := s.sr.RevokeAllSessionsByUserID(ctx, userID, now); err != nil {
			return fmt.Errorf("revoke all sessions by user id %s: %w", userID, err)
		}

		return nil
	}

	if err := s.sr.RevokeAllSessionByUserIDExceptCurrent(ctx, userID, currentSessionID, now); err != nil {
		return fmt.Errorf("revoke all sessions by user id %s except current %s: %w", userID, currentSessionID, err)
	}

	return nil
}
