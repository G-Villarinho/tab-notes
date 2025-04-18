package models

import "time"

type Like struct {
	UserID    string
	PostID    string
	CreatedAt time.Time
}
