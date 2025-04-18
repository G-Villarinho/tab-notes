package models

import (
	"errors"
	"time"
)

var (
	ErrCannotFollowSelf   = errors.New("cannot follow yourself")
	ErrCannotUnfollowSelf = errors.New("cannot unfollow yourself")
)

type Follower struct {
	UserID     string
	FollowerID string
	CreatedAt  time.Time
}

type FollowStats struct {
	Followers    int
	Following    int
	FollowedByMe bool
	FollowingMe  bool
}

type FollowerResponse struct {
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
