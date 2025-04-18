package services

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFollowUser(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if user not found", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.On("GetUserByUsername", ctx, "alice").Return(nil, nil)

		err := fs.FollowUser(ctx, "123", "alice")

		assert.ErrorIs(t, err, models.ErrUserNotFound)
		ur.AssertExpectations(t)
	})

	t.Run("should return error if trying to follow self", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.On("GetUserByUsername", ctx, "alice").Return(&models.User{ID: "123"}, nil)

		err := fs.FollowUser(ctx, "123", "alice")

		assert.ErrorIs(t, err, models.ErrCannotFollowSelf)
		ur.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.On("GetUserByUsername", ctx, "alice").Return(&models.User{ID: "999"}, nil)
		fr.On("CreateFollower", ctx, mock.AnythingOfType("*models.Follower")).Return(errors.New("repo fail"))

		err := fs.FollowUser(ctx, "123", "alice")

		assert.ErrorContains(t, err, "create follower")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should follow user successfully", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.On("GetUserByUsername", ctx, "alice").Return(&models.User{ID: "999"}, nil)
		fr.On("CreateFollower", ctx, mock.MatchedBy(func(f *models.Follower) bool {
			return f.UserID == "999" && f.FollowerID == "123"
		})).Return(nil)

		err := fs.FollowUser(ctx, "123", "alice")

		assert.NoError(t, err)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})
}

func TestUnfollowUser(t *testing.T) {
	ctx := context.Background()

	t.Run("should return ErrUserNotFound if user does not exist", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(nil, nil)

		err := fs.UnfollowUser(ctx, "user-id", "joaodasilva")

		assert.ErrorIs(t, err, models.ErrUserNotFound)
		ur.AssertExpectations(t)
	})

	t.Run("should return ErrCannotUnfollowSelf if trying to unfollow self", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(&models.User{ID: "user-id"}, nil)

		err := fs.UnfollowUser(ctx, "user-id", "joaodasilva")

		assert.ErrorIs(t, err, models.ErrCannotUnfollowSelf)
		ur.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(&models.User{ID: "user-123"}, nil)

		fr.
			On("DeleteFollower", ctx, "user-123", "user-456").
			Return(fmt.Errorf("delete error"))

		err := fs.UnfollowUser(ctx, "user-456", "joaodasilva")

		assert.ErrorContains(t, err, "delete follower")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should unfollow user successfully", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(&models.User{ID: "user-123"}, nil)

		fr.
			On("DeleteFollower", ctx, "user-123", "user-456").
			Return(nil)

		err := fs.UnfollowUser(ctx, "user-456", "joaodasilva")

		assert.NoError(t, err)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})
}

func TestGetFollowers(t *testing.T) {
	ctx := context.Background()

	t.Run("should return ErrUserNotFound if user does not exist", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(nil, nil)

		result, err := fs.GetFollowers(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		ur.AssertExpectations(t)
	})

	t.Run("should return error if GetUserByUsername fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetFollowers(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get user by username")
		ur.AssertExpectations(t)
	})

	t.Run("should return empty array if no followers", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowers", ctx, "u123").
			Return([]*models.Follower{}, nil)

		result, err := fs.GetFollowers(ctx, "joao")

		assert.NoError(t, err)
		assert.Empty(t, result)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetFollowers fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowers", ctx, "u123").
			Return(nil, fmt.Errorf("db fail"))

		result, err := fs.GetFollowers(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get followers IDs")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetUsersByIds fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowers", ctx, "u123").
			Return([]*models.Follower{{FollowerID: "f1"}}, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return(nil, fmt.Errorf("user fetch error"))

		result, err := fs.GetFollowers(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get users by IDs")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return followers when all succeeds", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowers", ctx, "u123").
			Return([]*models.Follower{{FollowerID: "f1"}}, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return([]*models.User{{Name: "Maria", Username: "maria"}}, nil)

		result, err := fs.GetFollowers(ctx, "joao")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "Maria", result[0].Name)
		assert.Equal(t, "maria", result[0].Username)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})
}

func TestGetFollowing(t *testing.T) {
	ctx := context.Background()

	t.Run("should return ErrUserNotFound if user does not exist", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(nil, nil)

		result, err := fs.GetFollowing(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		ur.AssertExpectations(t)
	})

	t.Run("should return error if GetUserByUsername fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetFollowing(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get user by username")
		ur.AssertExpectations(t)
	})

	t.Run("should return empty array if not following anyone", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowing", ctx, "u123").
			Return([]*models.Follower{}, nil)

		result, err := fs.GetFollowing(ctx, "joao")

		assert.NoError(t, err)
		assert.Empty(t, result)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetFollowing fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowing", ctx, "u123").
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetFollowing(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get following IDs")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetUsersByIds fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowing", ctx, "u123").
			Return([]*models.Follower{{UserID: "f1"}}, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return(nil, fmt.Errorf("user fetch fail"))

		result, err := fs.GetFollowing(ctx, "joao")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get users by IDs")
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})

	t.Run("should return following users successfully", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		ur.
			On("GetUserByUsername", ctx, "joao").
			Return(&models.User{ID: "u123"}, nil)

		fr.
			On("GetFollowing", ctx, "u123").
			Return([]*models.Follower{{UserID: "f1"}}, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return([]*models.User{{Name: "Ana", Username: "ana"}}, nil)

		result, err := fs.GetFollowing(ctx, "joao")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "Ana", result[0].Name)
		assert.Equal(t, "ana", result[0].Username)
		ur.AssertExpectations(t)
		fr.AssertExpectations(t)
	})
}

func TestGetMyFollowers(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if GetFollowers fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowers", ctx, "user-123").
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetMyFollowers(ctx, "user-123")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get followers IDs")
		fr.AssertExpectations(t)
	})

	t.Run("should return empty list if no followers", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowers", ctx, "user-123").
			Return([]*models.Follower{}, nil)

		result, err := fs.GetMyFollowers(ctx, "user-123")

		assert.NoError(t, err)
		assert.Empty(t, result)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetUsersByIds fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		followers := []*models.Follower{
			{FollowerID: "f1", CreatedAt: time.Now()},
		}

		fr.
			On("GetFollowers", ctx, "user-123").
			Return(followers, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetMyFollowers(ctx, "user-123")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get users by IDs")
		fr.AssertExpectations(t)
		ur.AssertExpectations(t)
	})

	t.Run("should return followers with createdAt", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		createdAt := time.Now()
		followers := []*models.Follower{
			{FollowerID: "f1", CreatedAt: createdAt},
		}

		users := []*models.User{
			{Name: "Ana", Username: "ana"},
		}

		fr.
			On("GetFollowers", ctx, "user-123").
			Return(followers, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"f1"}).
			Return(users, nil)

		result, err := fs.GetMyFollowers(ctx, "user-123")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "Ana", result[0].Name)
		assert.Equal(t, "ana", result[0].Username)
		assert.Equal(t, createdAt, *result[0].CreatedAt)

		fr.AssertExpectations(t)
		ur.AssertExpectations(t)
	})
}

func TestGetMyFollowing(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if GetFollowing fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowing", ctx, "user-123").
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetMyFollowing(ctx, "user-123")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get following IDs")
		fr.AssertExpectations(t)
	})

	t.Run("should return empty list if not following anyone", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowing", ctx, "user-123").
			Return([]*models.Follower{}, nil)

		result, err := fs.GetMyFollowing(ctx, "user-123")

		assert.NoError(t, err)
		assert.Empty(t, result)
		fr.AssertExpectations(t)
	})

	t.Run("should return error if GetUsersByIds fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		following := []*models.Follower{
			{UserID: "u1", CreatedAt: time.Now()},
		}

		fr.
			On("GetFollowing", ctx, "user-123").
			Return(following, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"u1"}).
			Return(nil, fmt.Errorf("db error"))

		result, err := fs.GetMyFollowing(ctx, "user-123")

		assert.Nil(t, result)
		assert.ErrorContains(t, err, "get users by IDs")
		fr.AssertExpectations(t)
		ur.AssertExpectations(t)
	})

	t.Run("should return following users with createdAt", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		createdAt := time.Now()
		following := []*models.Follower{
			{UserID: "u1", CreatedAt: createdAt},
		}

		users := []*models.User{
			{Name: "Carlos", Username: "carlos"},
		}

		fr.
			On("GetFollowing", ctx, "user-123").
			Return(following, nil)

		ur.
			On("GetUsersByIds", ctx, []string{"u1"}).
			Return(users, nil)

		result, err := fs.GetMyFollowing(ctx, "user-123")

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "Carlos", result[0].Name)
		assert.Equal(t, "carlos", result[0].Username)
		assert.Equal(t, createdAt, *result[0].CreatedAt)

		fr.AssertExpectations(t)
		ur.AssertExpectations(t)
	})
}

func TestGetFollowStats(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowStats", ctx, "user-123", "user-124").
			Return(nil, fmt.Errorf("query error"))

		stats, err := fs.GetFollowStats(ctx, "user-123", "user-124")

		assert.Nil(t, stats)
		assert.ErrorContains(t, err, "get follow stats")
		fr.AssertExpectations(t)
	})

	t.Run("should return ErrUserNotFound if result is nil", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		fr.
			On("GetFollowStats", ctx, "user-123", "user-124").
			Return(nil, nil)

		stats, err := fs.GetFollowStats(ctx, "user-123", "user-124")

		assert.Nil(t, stats)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		fr.AssertExpectations(t)
	})

	t.Run("should return follow stats successfully", func(t *testing.T) {
		fr := new(mocks.FollowerRepositoryMock)
		ur := new(mocks.UserRepositoryMock)
		fs := NewFollowerService(fr, ur)

		expected := &models.FollowStats{
			Followers:    10,
			Following:    5,
			FollowedByMe: true,
			FollowingMe:  false,
		}

		fr.
			On("GetFollowStats", ctx, "user-123", "user-124").
			Return(expected, nil)

		stats, err := fs.GetFollowStats(ctx, "user-123", "user-124")

		assert.NoError(t, err)
		assert.Equal(t, expected, stats)
		fr.AssertExpectations(t)
	})
}
