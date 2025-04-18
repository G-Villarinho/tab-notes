package services

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, name string, username string, email string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetProfile(ctx context.Context, id string) (*models.UserResponse, error)
	SearchUsers(ctx context.Context, query string) ([]*models.SearchUserResponse, error)
	GetProfileByUsername(ctx context.Context, username string, viewerID string) (*models.UserProfileResponse, error)
	UpdateUser(ctx context.Context, id string, name string, username string) error
}

type userService struct {
	fs FollowerService
	ur repositories.UserRepository
}

func NewUserService(
	followerService FollowerService,
	userRepository repositories.UserRepository) UserService {
	return &userService{
		ur: userRepository,
		fs: followerService,
	}
}

func (u *userService) CreateUser(ctx context.Context, name string, username string, email string) (*models.User, error) {
	userFromEmail, err := u.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	if userFromEmail != nil {
		return nil, models.ErrEmailAlreadyExists
	}

	userFromUsername, err := u.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	if userFromUsername != nil {
		return nil, models.ErrUsernameAlreadyExists
	}

	user := models.User{
		Name:      name,
		Email:     email,
		Username:  username,
		Status:    models.UserStatusActive,
		CreatedAt: time.Now().UTC(),
	}

	if err := u.ur.CreateUser(ctx, &user); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	return user, nil
}

func (u *userService) GetProfile(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := u.ur.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	followStats, err := u.fs.GetFollowStats(ctx, user.ID, "")
	if err != nil {
		return nil, fmt.Errorf("get follow stats: %w", err)
	}

	return &models.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Followers: followStats.Followers,
		Following: followStats.Following,
	}, nil
}

func (u *userService) SearchUsers(ctx context.Context, query string) ([]*models.SearchUserResponse, error) {
	users, err := u.ur.SearchUsers(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("search users: %w", err)
	}

	if len(users) == 0 {
		return []*models.SearchUserResponse{}, nil
	}

	var searchUsers []*models.SearchUserResponse
	for _, user := range users {
		searchUsers = append(searchUsers, &models.SearchUserResponse{
			Name:     user.Name,
			Username: user.Username,
		})
	}

	return searchUsers, nil
}

func (u *userService) GetProfileByUsername(ctx context.Context, username string, viewerID string) (*models.UserProfileResponse, error) {
	user, err := u.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	followStats, err := u.fs.GetFollowStats(ctx, user.ID, viewerID)
	if err != nil {
		return nil, fmt.Errorf("get follow stats: %w", err)
	}

	return &models.UserProfileResponse{
		Name:         user.Name,
		Username:     user.Username,
		Followers:    followStats.Followers,
		Following:    followStats.Following,
		FollowedByMe: followStats.FollowedByMe,
		FollowingMe:  followStats.FollowingMe,
	}, nil
}

func (u *userService) UpdateUser(ctx context.Context, id string, name string, username string) error {
	user, err := u.ur.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("get user by id %s: %w", id, err)
	}

	if user == nil {
		return models.ErrUserNotFound
	}

	if user.Username != username {
		userFromUsername, err := u.ur.GetUserByUsername(ctx, username)
		if err != nil {
			return fmt.Errorf("get user by username: %w", err)
		}

		if userFromUsername != nil {
			return models.ErrUsernameAlreadyExists
		}

		user.Username = username
	}

	user.Name = name
	if err := u.ur.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("update user: %w", err)
	}

	return nil
}
