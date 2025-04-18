package services

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type FollowerService interface {
	FollowUser(ctx context.Context, followerID string, username string) error
	UnfollowUser(ctx context.Context, followerID string, username string) error
	GetFollowers(ctx context.Context, username string) ([]*models.FollowerResponse, error)
	GetFollowing(ctx context.Context, username string) ([]*models.FollowerResponse, error)
	GetMyFollowers(ctx context.Context, userID string) ([]*models.FollowerResponse, error)
	GetMyFollowing(ctx context.Context, userID string) ([]*models.FollowerResponse, error)
	GetFollowStats(ctx context.Context, userID string, viewerID string) (*models.FollowStats, error)
}

type followerService struct {
	fr repositories.FollowerRepository
	ur repositories.UserRepository
}

func NewFollowerService(followerRepository repositories.FollowerRepository, userRepository repositories.UserRepository) FollowerService {
	return &followerService{
		fr: followerRepository,
		ur: userRepository,
	}
}

func (f *followerService) FollowUser(ctx context.Context, followerID string, username string) error {
	user, err := f.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("get user by username: %w", err)
	}

	if user == nil {
		return models.ErrUserNotFound
	}

	if user.ID == followerID {
		return models.ErrCannotFollowSelf
	}

	follower := &models.Follower{
		UserID:     user.ID,
		FollowerID: followerID,
		CreatedAt:  time.Now().UTC(),
	}

	if err := f.fr.CreateFollower(ctx, follower); err != nil {
		return fmt.Errorf("create follower: %w", err)
	}

	return nil
}

func (f *followerService) UnfollowUser(ctx context.Context, followerID string, username string) error {
	user, err := f.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return fmt.Errorf("get user by username: %w", err)
	}

	if user == nil {
		return models.ErrUserNotFound
	}

	if user.ID == followerID {
		return models.ErrCannotUnfollowSelf
	}

	if err := f.fr.DeleteFollower(ctx, user.ID, followerID); err != nil {
		return fmt.Errorf("delete follower: %w", err)
	}

	return nil
}

func (f *followerService) GetFollowers(ctx context.Context, username string) ([]*models.FollowerResponse, error) {
	user, err := f.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	followers, err := f.fr.GetFollowers(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("get followers IDs: %w", err)
	}

	if len(followers) == 0 {
		return []*models.FollowerResponse{}, nil
	}

	followersIDs := make([]string, len(followers))
	for i, follower := range followers {
		followersIDs[i] = follower.FollowerID
	}

	users, err := f.ur.GetUsersByIds(ctx, followersIDs)
	if err != nil {
		return nil, fmt.Errorf("get users by IDs: %w", err)
	}

	followersResponse := make([]*models.FollowerResponse, len(users))
	for i, u := range users {
		followersResponse[i] = &models.FollowerResponse{
			Name:     u.Name,
			Username: u.Username,
		}
	}

	return followersResponse, nil
}

func (f *followerService) GetFollowing(ctx context.Context, username string) ([]*models.FollowerResponse, error) {
	user, err := f.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	if user == nil {
		return nil, models.ErrUserNotFound
	}

	following, err := f.fr.GetFollowing(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("get following IDs: %w", err)
	}

	if len(following) == 0 {
		return []*models.FollowerResponse{}, nil
	}

	followingIDs := make([]string, len(following))
	for i, f := range following {
		followingIDs[i] = f.UserID
	}

	users, err := f.ur.GetUsersByIds(ctx, followingIDs)
	if err != nil {
		return nil, fmt.Errorf("get users by IDs: %w", err)
	}

	followingResponse := make([]*models.FollowerResponse, len(users))
	for i, user := range users {
		followingResponse[i] = &models.FollowerResponse{
			Name:     user.Name,
			Username: user.Username,
		}
	}

	return followingResponse, nil
}

func (f *followerService) GetMyFollowers(ctx context.Context, userID string) ([]*models.FollowerResponse, error) {
	followers, err := f.fr.GetFollowers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get followers IDs: %w", err)
	}

	if len(followers) == 0 {
		return []*models.FollowerResponse{}, nil
	}

	followersIDs := make([]string, len(followers))
	for i, follower := range followers {
		followersIDs[i] = follower.FollowerID
	}

	users, err := f.ur.GetUsersByIds(ctx, followersIDs)
	if err != nil {
		return nil, fmt.Errorf("get users by IDs: %w", err)
	}

	followersResponse := make([]*models.FollowerResponse, len(users))
	for i, user := range users {
		followersResponse[i] = &models.FollowerResponse{
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: &followers[i].CreatedAt,
		}
	}

	return followersResponse, nil
}

func (f *followerService) GetMyFollowing(ctx context.Context, userID string) ([]*models.FollowerResponse, error) {
	following, err := f.fr.GetFollowing(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get following IDs: %w", err)
	}

	if len(following) == 0 {
		return []*models.FollowerResponse{}, nil
	}

	followingIDs := make([]string, len(following))
	for i, f := range following {
		followingIDs[i] = f.UserID
	}

	users, err := f.ur.GetUsersByIds(ctx, followingIDs)
	if err != nil {
		return nil, fmt.Errorf("get users by IDs: %w", err)
	}

	followingResponse := make([]*models.FollowerResponse, len(users))
	for i, user := range users {
		followingResponse[i] = &models.FollowerResponse{
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: &following[i].CreatedAt,
		}
	}

	return followingResponse, nil
}

func (f *followerService) GetFollowStats(ctx context.Context, userID string, viewerID string) (*models.FollowStats, error) {
	followStats, err := f.fr.GetFollowStats(ctx, userID, viewerID)
	if err != nil {
		return nil, fmt.Errorf("get follow stats: %w", err)
	}

	if followStats == nil {
		return nil, models.ErrUserNotFound
	}

	return followStats, nil
}
