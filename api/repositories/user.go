package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/utils"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUsersByIds(ctx context.Context, ids []string) ([]*models.User, error)
	SearchUsers(ctx context.Context, query string) ([]*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user.ID = id.String()

	query := `
		INSERT INTO users (id, name, username, email, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Username, user.Email, user.Status, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, username, email, status, created_at, updated_at, banned_at
		FROM users
		WHERE email = ?
	`

	return utils.QueryRowScan(ctx, r.db, query, scanUser, email)
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, username, email, status, created_at, updated_at, banned_at
		FROM users
		WHERE id = ?
	`

	return utils.QueryRowScan(ctx, r.db, query, scanUser, id)
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
		SELECT id, name, username, email, status, created_at, updated_at, banned_at
		FROM users
		WHERE username = ?
	`

	return utils.QueryRowScan(ctx, r.db, query, scanUser, username)
}

func (r *userRepository) GetUsersByIds(ctx context.Context, ids []string) ([]*models.User, error) {
	if len(ids) == 0 {
		return []*models.User{}, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := `
		SELECT id, name, username, email, status, created_at, updated_at, banned_at
		FROM users
		WHERE id IN (` + join(placeholders, ",") + `)
	`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Email, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.BannedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) SearchUsers(ctx context.Context, query string) ([]*models.User, error) {
	search := "%" + query + "%"
	sqlQuery := `
		SELECT name, username
		FROM users
		WHERE name LIKE ? OR username LIKE ?
	`

	rows, err := r.db.QueryContext(ctx, sqlQuery, search, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Name, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	updatedAt := sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	sql := `
		UPDATE users
		SET name = ?, username = ?, updated_at = ?
		WHERE id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Username, updatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	result := strs[0]
	for _, s := range strs[1:] {
		result += sep + s
	}
	return result
}

func scanUser(row *sql.Row) (*models.User, error) {
	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Username, &u.Email, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.BannedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
