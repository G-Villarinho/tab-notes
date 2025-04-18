package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrPostNotFound        = errors.New("post not found")
	ErrPostNotBelongToUser = errors.New("post does not belong to user")
)

type Post struct {
	ID        string
	Title     string
	Content   string
	AuthorID  string
	Likes     int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Likes       int       `json:"likes"`
	LikedByUser bool      `json:"liked_by_user"`
	CreatedAt   time.Time `json:"created_at"`
}
