package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	t.Run("should return ErrEmailALreadyExists if user already exists by email", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(&models.User{}, nil)

		user, err := userService.CreateUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.Nil(t, user)
		assert.ErrorIs(t, err, models.ErrEmailAlreadyExists)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return ErrUsernaeAlreadyExists if user already exists by username", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, nil)

		userRepo.
			On("GetUserByUsername", ctx, "joaozinho-30").
			Return(&models.User{}, nil)

		user, err := userService.CreateUser(ctx, "João", "joaozinho-30", "joao@example.com")

		assert.Nil(t, user)
		assert.ErrorIs(t, err, models.ErrUsernameAlreadyExists)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error if get by email fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, errors.New("db error"))

		user, err := userService.CreateUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.Nil(t, user)
		assert.ErrorContains(t, err, "get user by email")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error if get by username fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, nil)

		userRepo.
			On("GetUserByUsername", ctx, "joaozinha-30").
			Return(nil, errors.New("db username error"))

		user, err := userService.CreateUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.Nil(t, user)
		assert.ErrorContains(t, err, "get user by username")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error if user creation fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, nil)

		userRepo.
			On("GetUserByUsername", ctx, "joaozinha-30").
			Return(nil, nil)

		userRepo.
			On("CreateUser", ctx, mock.AnythingOfType("*models.User")).
			Return(errors.New("insert fail"))

		user, err := userService.CreateUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.Nil(t, user)
		assert.ErrorContains(t, err, "create user")
		userRepo.AssertExpectations(t)
	})

	t.Run("should create user successfully", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(nil, nil)

		userRepo.
			On("GetUserByUsername", ctx, "joaozinha-30").
			Return(nil, nil)

		userRepo.
			On("CreateUser", ctx, mock.AnythingOfType("*models.User")).
			Return(nil)

		user, err := userService.CreateUser(ctx, "João", "joaozinha-30", "joao@example.com")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "João", user.Name)
		assert.Equal(t, "joao@example.com", user.Email)
		assert.Equal(t, models.UserStatusActive, user.Status)
		assert.WithinDuration(t, time.Now(), user.CreatedAt, time.Second)
		userRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmail(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "fail@example.com").
			Return(nil, errors.New("db error"))

		user, err := userService.GetUserByEmail(ctx, "fail@example.com")

		assert.Nil(t, user)
		assert.ErrorContains(t, err, "get user by email")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return ErrUserNotFound if user is nil", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByEmail", ctx, "notfound@example.com").
			Return(nil, nil)

		user, err := userService.GetUserByEmail(ctx, "notfound@example.com")

		assert.Nil(t, user)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return user when found", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		expected := &models.User{
			ID:    "123",
			Name:  "João",
			Email: "joao@example.com",
		}

		userRepo.
			On("GetUserByEmail", ctx, "joao@example.com").
			Return(expected, nil)

		user, err := userService.GetUserByEmail(ctx, "joao@example.com")

		assert.NoError(t, err)
		assert.Equal(t, expected, user)
		userRepo.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repo fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByID", ctx, "123").
			Return(nil, errors.New("db error"))

		profile, err := userService.GetProfile(ctx, "123")

		assert.Nil(t, profile)
		assert.ErrorContains(t, err, "get user by id")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return ErrUserNotFound if user is nil", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByID", ctx, "notfound").
			Return(nil, nil)

		profile, err := userService.GetProfile(ctx, "notfound")

		assert.Nil(t, profile)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error if follow stats fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		mockUser := &models.User{
			ID:     "abc123",
			Name:   "João da Silva",
			Email:  "joao@example.com",
			Status: models.UserStatusActive,
		}

		userRepo.
			On("GetUserByID", ctx, "abc123").
			Return(mockUser, nil)

		followerService.
			On("GetFollowStats", ctx, "abc123", "").
			Return((*models.FollowStats)(nil), errors.New("failed stats"))

		profile, err := userService.GetProfile(ctx, "abc123")

		assert.Nil(t, profile)
		assert.ErrorContains(t, err, "get follow stats")

		userRepo.AssertExpectations(t)
		followerService.AssertExpectations(t)
	})
	t.Run("should return profile when user is found", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		mockUser := &models.User{
			ID:       "abc123",
			Name:     "João da Silva",
			Username: "joaosilva",
			Email:    "joao@example.com",
			Status:   models.UserStatusActive,
		}

		mockStats := &models.FollowStats{
			Followers: 10,
			Following: 5,
		}

		userRepo.
			On("GetUserByID", ctx, "abc123").
			Return(mockUser, nil)

		followerService.
			On("GetFollowStats", ctx, "abc123", "").
			Return(mockStats, nil)

		profile, err := userService.GetProfile(ctx, "abc123")

		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, mockUser.Name, profile.Name)
		assert.Equal(t, mockUser.Email, profile.Email)
		assert.Equal(t, mockUser.Username, profile.Username)
		assert.Equal(t, mockStats.Followers, profile.Followers)
		assert.Equal(t, mockStats.Following, profile.Following)

		userRepo.AssertExpectations(t)
		followerService.AssertExpectations(t)
	})
}

func TestSearchUsers(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repo fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("SearchUsers", ctx, "joao").
			Return(nil, errors.New("db error"))

		users, err := userService.SearchUsers(ctx, "joao")

		assert.Nil(t, users)
		assert.ErrorContains(t, err, "search users")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return empty array if no users are found", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("SearchUsers", ctx, "joao").
			Return([]*models.User{}, nil)

		users, err := userService.SearchUsers(ctx, "joao")

		assert.NoError(t, err)
		assert.Empty(t, users)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return users", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		expected := []*models.User{
			{Name: "João da Silva", Username: "joaodasilva"},
			{Name: "João Pedro", Username: "joaopedro"},
		}

		userRepo.
			On("SearchUsers", ctx, "joao").
			Return(expected, nil)

		users, err := userService.SearchUsers(ctx, "joao")

		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, expected[0].Name, users[0].Name)
		assert.Equal(t, expected[1].Username, users[1].Username)
		userRepo.AssertExpectations(t)
	})
}

func TestGetProfileByUsername(t *testing.T) {
	ctx := context.Background()

	t.Run("should return ErrUserNotFound if user doesn't exist", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(nil, nil)

		resp, err := userService.GetProfileByUsername(ctx, "joaodasilva", "123")

		assert.Nil(t, resp)
		assert.ErrorIs(t, err, models.ErrUserNotFound)
		userRepo.AssertExpectations(t)
	})

	t.Run("should return profile data if user exists", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		mockUser := &models.User{
			ID:       "user-123",
			Name:     "João da Silva",
			Username: "joaodasilva",
		}

		mockStats := &models.FollowStats{
			Followers:    42,
			Following:    17,
			FollowedByMe: true,
			FollowingMe:  false,
		}

		userRepo.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(mockUser, nil)

		followerService.
			On("GetFollowStats", ctx, "user-123", "123").
			Return(mockStats, nil)

		resp, err := userService.GetProfileByUsername(ctx, "joaodasilva", "123")

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, resp.Name)
		assert.Equal(t, mockUser.Username, resp.Username)
		assert.Equal(t, mockStats.Followers, resp.Followers)
		assert.Equal(t, mockStats.Following, resp.Following)
		assert.Equal(t, mockStats.FollowedByMe, resp.FollowedByMe)
		assert.Equal(t, mockStats.FollowingMe, resp.FollowingMe)

		userRepo.AssertExpectations(t)
		followerService.AssertExpectations(t)
	})

	t.Run("should return error if userRepo fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		userRepo.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(nil, errors.New("db error"))

		resp, err := userService.GetProfileByUsername(ctx, "joaodasilva", "123")

		assert.Nil(t, resp)
		assert.ErrorContains(t, err, "get user by username")
		userRepo.AssertExpectations(t)
	})

	t.Run("should return error if followerService fails", func(t *testing.T) {
		userRepo := new(mocks.UserRepositoryMock)
		followerService := new(mocks.FollowerServiceMock)
		userService := NewUserService(followerService, userRepo)

		mockUser := &models.User{
			ID:       "user-123",
			Name:     "João da Silva",
			Username: "joaodasilva",
		}

		userRepo.
			On("GetUserByUsername", ctx, "joaodasilva").
			Return(mockUser, nil)

		followerService.
			On("GetFollowStats", ctx, "user-123", "user-123").
			Return(nil, errors.New("stats error"))

		resp, err := userService.GetProfileByUsername(ctx, "joaodasilva", "user-123")

		assert.Nil(t, resp)
		assert.ErrorContains(t, err, "get follow stats")

		userRepo.AssertExpectations(t)
		followerService.AssertExpectations(t)
	})
}
