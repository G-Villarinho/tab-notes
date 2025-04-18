package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/models"
)

type FeedRepository interface {
	GetFeed(ctx context.Context, userID string, limit, offset int) ([]*models.FeedPostResponse, error)
}

type feedRepository struct {
	db *sql.DB
}

func NewFeedRepository(db *sql.DB) FeedRepository {
	return &feedRepository{
		db: db,
	}
}

func (r *feedRepository) GetFeed(ctx context.Context, userID string, limit, offset int) ([]*models.FeedPostResponse, error) {
	query := `
		SELECT p.id, p.title, p.content, p.likes, p.created_at,
		       u.name AS author_name, u.username AS author_username
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN followers f ON f.user_id = p.author_id AND f.follower_id = ?
		WHERE f.follower_id IS NOT NULL OR p.author_id = ?
		ORDER BY p.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.QueryContext(ctx, query, userID, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query feed: %w", err)
	}
	defer rows.Close()

	var feed []*models.FeedPostResponse
	for rows.Next() {
		var post models.FeedPostResponse
		err := rows.Scan(&post.PostID, &post.Title, &post.Content, &post.Likes, &post.CreatedAt, &post.AuthorName, &post.AuthorUsername)
		if err != nil {
			return nil, fmt.Errorf("scan feed post: %w", err)
		}
		feed = append(feed, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return feed, nil
}
