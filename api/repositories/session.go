package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/google/uuid"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session *models.Session) error
	GetSessionByToken(ctx context.Context, token string) (*models.Session, error)
	DeleteSession(ctx context.Context, id string) error
	UpdateSession(ctx context.Context, session *models.Session) error
	GetSessionsByUserID(ctx context.Context, userID string) ([]*models.Session, error)
	RevokeSession(ctx context.Context, id string, revokedAt time.Time) error
	IsSessionRevoked(ctx context.Context, id string) (bool, error)
	GetSessionById(ctx context.Context, id string) (*models.Session, error)
	RevokeAllSessionsByUserID(ctx context.Context, userID string, revoketAt time.Time) error
	RevokeAllSessionByUserIDExceptCurrent(ctx context.Context, userID string, currentSessionID string, revokedAt time.Time) error
}

type sessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) SessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (r *sessionRepository) CreateSession(ctx context.Context, session *models.Session) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("generate UUID: %w", err)
	}
	session.ID = id.String()

	query := `
		INSERT INTO sessions (id, token, expires_at, user_id, created_at)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, session.ID, session.Token, session.ExpiresAt, session.UserID, session.CreatedAt)
	if err != nil {
		return err
	}

	return err
}

func (r *sessionRepository) GetSessionByToken(ctx context.Context, token string) (*models.Session, error) {
	query := `
		SELECT id, token, expires_at, user_id, revoked_at, verified_at, created_at, updated_at
		FROM sessions
		WHERE token = ?
	`

	row := r.db.QueryRowContext(ctx, query, token)

	var session models.Session
	err := row.Scan(
		&session.ID,
		&session.Token,
		&session.ExpiresAt,
		&session.UserID,
		&session.RevokedAt,
		&session.VerifiedAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *sessionRepository) UpdateSession(ctx context.Context, session *models.Session) error {
	session.UpdatedAt = sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	query := `
		UPDATE sessions
		SET token = ?, expires_at = ?, user_id = ?, revoked_at = ?, verified_at = ?, updated_at = ?
		WHERE id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, session.Token, session.ExpiresAt, session.UserID, session.RevokedAt, session.VerifiedAt, session.UpdatedAt, session.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, id string) error {
	query := `
		DELETE FROM sessions
		WHERE id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionRepository) GetSessionsByUserID(ctx context.Context, userID string) ([]*models.Session, error) {
	query := `
		SELECT id, token, expires_at, user_id, revoked_at, verified_at, created_at, updated_at
		FROM sessions
		WHERE user_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*models.Session
	for rows.Next() {
		var s models.Session
		if err := rows.Scan(&s.ID, &s.Token, &s.ExpiresAt, &s.UserID, &s.RevokedAt, &s.VerifiedAt, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		sessions = append(sessions, &s)
	}

	return sessions, nil
}

func (r *sessionRepository) RevokeSession(ctx context.Context, id string, revokedAt time.Time) error {
	query := `UPDATE sessions SET revoked_at = ?, updated_at = ? WHERE id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, revokedAt, sql.NullTime{Time: time.Now().UTC(), Valid: true}, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionRepository) IsSessionRevoked(ctx context.Context, id string) (bool, error) {
	var revokedAt sql.NullTime
	err := r.db.QueryRowContext(ctx, "SELECT revoked_at FROM sessions WHERE id = ?", id).Scan(&revokedAt)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return revokedAt.Valid, nil
}

func (r *sessionRepository) GetSessionById(ctx context.Context, id string) (*models.Session, error) {
	query := `
		SELECT id, token, expires_at, user_id, revoked_at, verified_at, created_at, updated_at
		FROM sessions
		WHERE id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var session models.Session
	err := row.Scan(
		&session.ID,
		&session.Token,
		&session.ExpiresAt,
		&session.UserID,
		&session.RevokedAt,
		&session.VerifiedAt,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *sessionRepository) RevokeAllSessionsByUserID(ctx context.Context, userID string, revoketAt time.Time) error {
	query := `UPDATE sessions SET revoked_at = ?, updated_at = ? WHERE user_id = ?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, revoketAt, sql.NullTime{Time: time.Now().UTC(), Valid: true}, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *sessionRepository) RevokeAllSessionByUserIDExceptCurrent(ctx context.Context, userID string, currentSessionID string, revokedAt time.Time) error {
	query := `UPDATE sessions SET revoked_at = ?, updated_at = ? WHERE user_id = ? AND id != ?`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, revokedAt, sql.NullTime{Time: time.Now().UTC(), Valid: true}, userID, currentSessionID)
	if err != nil {
		return err
	}

	return nil
}
