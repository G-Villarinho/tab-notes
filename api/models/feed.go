package models

import "time"

type FeedPostResponse struct {
	PostID         string    `json:"post_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Likes          int       `json:"likes"`
	CreatedAt      time.Time `json:"created_at"`
	AuthorName     string    `json:"author_name"`
	AuthorUsername string    `json:"author_username"`
	LikedByUser    bool      `json:"liked_by_user"`
}
