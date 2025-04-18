package repositories

import (
	"context"
	"database/sql"

	"github.com/g-villarinho/tab-notes-api/models"
)

type FollowerRepository interface {
	CreateFollower(ctx context.Context, follower *models.Follower) error
	DeleteFollower(ctx context.Context, userID string, followerID string) error
	GetFollowers(ctx context.Context, userID string) ([]*models.Follower, error)
	GetFollowing(ctx context.Context, followerID string) ([]*models.Follower, error)
	CountFollowers(ctx context.Context, userID string) (int, error)
	CountFollowing(ctx context.Context, userID string) (int, error)
	GetFollowStats(ctx context.Context, userID, viewerID string) (*models.FollowStats, error)
}

type followerRepository struct {
	db *sql.DB
}

func NewFollowerRepository(db *sql.DB) FollowerRepository {
	return &followerRepository{
		db: db,
	}
}

func (f *followerRepository) CreateFollower(ctx context.Context, follower *models.Follower) error {
	query := `
		INSERT IGNORE INTO followers (user_id, follower_id, created_at)
		VALUES (?, ?, ?)
	`

	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, follower.UserID, follower.FollowerID, follower.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (f *followerRepository) DeleteFollower(ctx context.Context, userID string, followerID string) error {
	query := `
		DELETE FROM followers
		WHERE user_id = ? AND follower_id = ?
	`

	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (f *followerRepository) GetFollowers(ctx context.Context, userID string) ([]*models.Follower, error) {
	query := `
		SELECT follower_id, created_at
		FROM followers
		WHERE user_id = ?
	`

	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []*models.Follower
	for rows.Next() {
		follower := &models.Follower{}
		if err := rows.Scan(&follower.FollowerID, &follower.CreatedAt); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (f *followerRepository) GetFollowing(ctx context.Context, followerID string) ([]*models.Follower, error) {
	query := `
		SELECT user_id, created_at
		FROM followers
		WHERE follower_id = ?
	`

	stmt, err := f.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, followerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []*models.Follower
	for rows.Next() {
		follower := &models.Follower{}
		if err := rows.Scan(&follower.UserID, &follower.CreatedAt); err != nil {
			return nil, err
		}
		following = append(following, follower)
	}

	return following, nil
}

func (r *followerRepository) CountFollowers(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM followers WHERE user_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	return count, err
}

func (r *followerRepository) CountFollowing(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM followers WHERE follower_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	return count, err
}

func (f *followerRepository) GetFollowStats(ctx context.Context, userID string, viewerID string) (*models.FollowStats, error) {
	query := `
		SELECT
			(SELECT COUNT(*) FROM followers WHERE user_id = ?) AS followers,
			(SELECT COUNT(*) FROM followers WHERE follower_id = ?) AS following,
			EXISTS(SELECT 1 FROM followers WHERE user_id = ? AND follower_id = ?) AS followed_by_me,
			EXISTS(SELECT 1 FROM followers WHERE user_id = ? AND follower_id = ?) AS following_me
	`

	row := f.db.QueryRowContext(ctx, query,
		userID,
		userID,
		userID, viewerID,
		viewerID, userID,
	)

	var stats models.FollowStats
	if err := row.Scan(
		&stats.Followers,
		&stats.Following,
		&stats.FollowedByMe,
		&stats.FollowingMe,
	); err != nil {
		return nil, err
	}

	return &stats, nil
}
