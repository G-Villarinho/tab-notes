package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
)

type User struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	BannedAt  sql.NullTime
}

type CreateUserPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

type UserProfileResponse struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
	FollowedByMe bool   `json:"followed_by_me"`
	FollowingMe  bool   `json:"following_me"`
}

type SearchUserResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}
