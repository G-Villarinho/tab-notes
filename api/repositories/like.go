package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
)

type LikeRepository interface {
	CreateLike(ctx context.Context, like *models.Like) error
	DeleteLike(ctx context.Context, like *models.Like) error
	CheckLike(ctx context.Context, userID, postID string) (bool, error)
	GetLikedPostIDs(ctx context.Context, userID string, postIDs []string) ([]string, error)
}

type likeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) LikeRepository {
	return &likeRepository{
		db: db,
	}
}

func (l *likeRepository) CreateLike(ctx context.Context, like *models.Like) error {
	like.CreatedAt = time.Now().UTC()

	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	insertQuery := `
		INSERT IGNORE INTO likes (user_id, post_id, created_at)
		VALUES (?, ?, ?)
	`
	res, err := tx.ExecContext(ctx, insertQuery, like.UserID, like.PostID, like.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if rowsAffected > 0 {
		updateQuery := `
			UPDATE posts
			SET likes = likes + 1
			WHERE id = ?
		`
		_, err = tx.ExecContext(ctx, updateQuery, like.PostID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (l *likeRepository) DeleteLike(ctx context.Context, like *models.Like) error {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	deleteQuery := `
		DELETE FROM likes
		WHERE user_id = ? AND post_id = ?
	`
	res, err := tx.ExecContext(ctx, deleteQuery, like.UserID, like.PostID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if rowsAffected > 0 {
		updateQuery := `
			UPDATE posts
			SET likes = likes - 1
			WHERE id = ?
		`
		_, err = tx.ExecContext(ctx, updateQuery, like.PostID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (l *likeRepository) CheckLike(ctx context.Context, userID string, postID string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM likes
		WHERE user_id = ? AND post_id = ?
	`
	var exists bool
	if err := l.db.QueryRowContext(ctx, query, userID, postID).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (l *likeRepository) GetLikedPostIDs(ctx context.Context, userID string, postIDs []string) ([]string, error) {
	if len(postIDs) == 0 {
		return nil, nil
	}

	placeholders := strings.Repeat("?,", len(postIDs))
	placeholders = placeholders[:len(placeholders)-1]

	args := make([]any, 0, len(postIDs)+1)
	args = append(args, userID)
	for _, id := range postIDs {
		args = append(args, id)
	}

	query := fmt.Sprintf(`
		SELECT post_id FROM likes
		WHERE user_id = ? AND post_id IN (%s)
	`, placeholders)

	rows, err := l.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likedPostIDs []string
	for rows.Next() {
		var postID string
		if err := rows.Scan(&postID); err != nil {
			return nil, err
		}
		likedPostIDs = append(likedPostIDs, postID)
	}

	return likedPostIDs, nil
}
