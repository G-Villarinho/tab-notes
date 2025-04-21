package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrSessionNotFound        = errors.New("session not found")
	ErrSessionExpired         = errors.New("session expired")
	ErrSessionNotBelongToUser = errors.New("session does not belong to user")
)

type Session struct {
	ID         string
	Token      string
	ExpiresAt  time.Time
	UserID     string
	VerifiedAt sql.NullTime
	RevokedAt  sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type RevokeAllSessionsPayload struct {
	RevokeCurrent bool `json:"revoke_current"`
}

type SessionResponse struct {
	ID               string     `json:"id"`
	ExpiresAt        time.Time  `json:"expires_at"`
	CurrentSessionID string     `json:"current_session_id"`
	VerifiedAt       *time.Time `json:"verified_at,omitempty"`
	RevokedAt        *time.Time `json:"revoked_at,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
}
