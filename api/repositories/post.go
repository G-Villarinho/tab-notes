package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/google/uuid"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, ID string) (*models.Post, error)
	GetPostsByAuthorID(ctx context.Context, authorID string) ([]*models.Post, error)
	DeletePost(ctx context.Context, ID string) error
	UpdatePost(ctx context.Context, post *models.Post) error
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) CreatePost(ctx context.Context, post *models.Post) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	post.ID = id.String()
	post.CreatedAt = time.Now().UTC()

	query := `INSERT INTO posts (id, title, content, author_id, created_at) VALUES (?, ?, ?, ?, ?)`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, post.ID, post.Title, post.Content, post.AuthorID, post.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *postRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	query := `SELECT id, title, content, author_id, likes, created_at FROM posts WHERE id = ?`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	post := &models.Post{}
	err = row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return post, nil
}

func (p *postRepository) GetPostsByAuthorID(ctx context.Context, authorID string) ([]*models.Post, error) {
	query := `SELECT id, title, content, author_id, likes, created_at FROM posts WHERE author_id = ?`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		post := &models.Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *postRepository) DeletePost(ctx context.Context, ID string) error {
	query := `DELETE FROM posts WHERE id = ?`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *postRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	post.UpdatedAt = sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	query := `UPDATE posts SET title = ?, content = ?, likes = ?, updated_at = ? WHERE id = ?`

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, post.Title, post.Content, post.Likes, post.UpdatedAt, post.ID)
	if err != nil {
		return err
	}

	return nil
}
