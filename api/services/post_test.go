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

func TestCreatePost(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("CreatePost", ctx, mock.AnythingOfType("*models.Post")).
			Return(errors.New("db error"))

		_, err := ps.CreatePost(ctx, "user-123", "Meu título", "Meu conteúdo")

		assert.ErrorContains(t, err, "create post")
		postRepo.AssertExpectations(t)
	})

	t.Run("should create post successfully", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("CreatePost", ctx, mock.AnythingOfType("*models.Post")).
			Return(nil)

		_, err := ps.CreatePost(ctx, "user-123", "Título válido", "Conteúdo válido")

		assert.NoError(t, err)
		postRepo.AssertExpectations(t)
	})
}

func TestLikePost(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if GetPostByID fails", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(nil, errors.New("db error"))

		err := ps.LikePost(ctx, "user-123", "post-123")

		assert.ErrorContains(t, err, "get post by id")
		postRepo.AssertExpectations(t)
	})

	t.Run("should return ErrPostNotFound if post is nil", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(nil, nil)

		err := ps.LikePost(ctx, "user-123", "post-123")

		assert.ErrorIs(t, err, models.ErrPostNotFound)
		postRepo.AssertExpectations(t)
	})

	t.Run("should return error if LikePost fails", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		post := &models.Post{ID: "post-123"}

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(post, nil)

		likeService.
			On("LikePost", ctx, "user-123", "post-123").
			Return(errors.New("like failed"))

		err := ps.LikePost(ctx, "user-123", "post-123")

		assert.ErrorContains(t, err, "like post")
		postRepo.AssertExpectations(t)
		likeService.AssertExpectations(t)
	})

	t.Run("should like post successfully", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		post := &models.Post{ID: "post-123"}

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(post, nil)

		likeService.
			On("LikePost", ctx, "user-123", "post-123").
			Return(nil)

		err := ps.LikePost(ctx, "user-123", "post-123")

		assert.NoError(t, err)
		postRepo.AssertExpectations(t)
		likeService.AssertExpectations(t)
	})
}

func TestUnlikePost(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if GetPostByID fails", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(nil, errors.New("db error"))

		err := ps.UnlikePost(ctx, "user-123", "post-123")

		assert.ErrorContains(t, err, "get post by id")
		postRepo.AssertExpectations(t)
	})

	t.Run("should return ErrPostNotFound if post is nil", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(nil, nil)

		err := ps.UnlikePost(ctx, "user-123", "post-123")

		assert.ErrorIs(t, err, models.ErrPostNotFound)
		postRepo.AssertExpectations(t)
	})

	t.Run("should return error if UnlikePost fails", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		post := &models.Post{ID: "post-123"}

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(post, nil)

		likeService.
			On("UnlikePost", ctx, "user-123", "post-123").
			Return(errors.New("unlike failed"))

		err := ps.UnlikePost(ctx, "user-123", "post-123")

		assert.ErrorContains(t, err, "unlike post")
		postRepo.AssertExpectations(t)
		likeService.AssertExpectations(t)
	})

	t.Run("should unlike post successfully", func(t *testing.T) {
		likeService := new(mocks.LikeServiceMock)
		postRepo := new(mocks.PostRepositoryMock)
		userRepo := new(mocks.UserRepositoryMock)
		ps := NewPostService(likeService, postRepo, userRepo)

		post := &models.Post{ID: "post-123"}

		postRepo.
			On("GetPostByID", ctx, "post-123").
			Return(post, nil)

		likeService.
			On("UnlikePost", ctx, "user-123", "post-123").
			Return(nil)

		err := ps.UnlikePost(ctx, "user-123", "post-123")

		assert.NoError(t, err)
		postRepo.AssertExpectations(t)
		likeService.AssertExpectations(t)
	})
}

func TestPostService_GetPostByID(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ls := new(mocks.LikeServiceMock)
		ps := NewPostService(ls, pr, nil)

		pr.On("GetPostByID", ctx, "123").Return(nil, errors.New("db error"))

		post, err := ps.GetPostByID(ctx, "user1", "123")

		assert.Nil(t, post)
		assert.ErrorContains(t, err, "get post by id")
		pr.AssertExpectations(t)
	})

	t.Run("should return nil if post not found", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ls := new(mocks.LikeServiceMock)
		ps := NewPostService(ls, pr, nil)

		pr.On("GetPostByID", ctx, "123").Return(nil, nil)

		post, err := ps.GetPostByID(ctx, "user1", "123")

		assert.Nil(t, post)
		assert.NoError(t, err)
		pr.AssertExpectations(t)
	})

	t.Run("should return error if like check fails", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ls := new(mocks.LikeServiceMock)
		ps := NewPostService(ls, pr, nil)

		mockPost := &models.Post{
			ID:        "123",
			Title:     "Post",
			Content:   "Content",
			Likes:     10,
			CreatedAt: time.Now(),
		}

		pr.On("GetPostByID", ctx, "123").Return(mockPost, nil)
		ls.On("CheckLike", ctx, "user1", "123").Return(false, errors.New("check failed"))

		post, err := ps.GetPostByID(ctx, "user1", "123")

		assert.Nil(t, post)
		assert.ErrorContains(t, err, "check like")
		pr.AssertExpectations(t)
		ls.AssertExpectations(t)
	})

	t.Run("should return post response successfully", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ls := new(mocks.LikeServiceMock)
		ps := NewPostService(ls, pr, nil)

		mockPost := &models.Post{
			ID:        "123",
			Title:     "Post",
			Content:   "Content",
			Likes:     10,
			CreatedAt: time.Now(),
		}

		pr.On("GetPostByID", ctx, "123").Return(mockPost, nil)
		ls.On("CheckLike", ctx, "user1", "123").Return(true, nil)

		post, err := ps.GetPostByID(ctx, "user1", "123")

		assert.NoError(t, err)
		assert.NotNil(t, post)
		assert.Equal(t, mockPost.ID, post.ID)
		assert.True(t, post.LikedByUser)
		pr.AssertExpectations(t)
		ls.AssertExpectations(t)
	})
}

func TestPostService_DeletePost(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if get post fails", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ps := NewPostService(nil, pr, nil)

		pr.On("GetPostByID", ctx, "123").Return(nil, errors.New("db error"))

		err := ps.DeletePost(ctx, "user1", "123")

		assert.ErrorContains(t, err, "get post by id")
		pr.AssertExpectations(t)
	})

	t.Run("should return ErrPostNotFound if post is nil", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ps := NewPostService(nil, pr, nil)

		pr.On("GetPostByID", ctx, "123").Return(nil, nil)

		err := ps.DeletePost(ctx, "user1", "123")

		assert.ErrorIs(t, err, models.ErrPostNotFound)
		pr.AssertExpectations(t)
	})

	t.Run("should return ErrPostNotBelongToUser if user is not the author", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ps := NewPostService(nil, pr, nil)

		post := &models.Post{
			ID:       "123",
			AuthorID: "otherUser",
		}

		pr.On("GetPostByID", ctx, "123").Return(post, nil)

		err := ps.DeletePost(ctx, "user1", "123")

		assert.ErrorIs(t, err, models.ErrPostNotBelongToUser)
		pr.AssertExpectations(t)
	})

	t.Run("should return error if delete fails", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ps := NewPostService(nil, pr, nil)

		post := &models.Post{
			ID:       "123",
			AuthorID: "user1",
		}

		pr.On("GetPostByID", ctx, "123").Return(post, nil)
		pr.On("DeletePost", ctx, "123").Return(errors.New("delete error"))

		err := ps.DeletePost(ctx, "user1", "123")

		assert.ErrorContains(t, err, "delete post")
		pr.AssertExpectations(t)
	})

	t.Run("should delete post successfully", func(t *testing.T) {
		pr := new(mocks.PostRepositoryMock)
		ps := NewPostService(nil, pr, nil)

		post := &models.Post{
			ID:       "123",
			AuthorID: "user1",
		}

		pr.On("GetPostByID", ctx, "123").Return(post, nil)
		pr.On("DeletePost", ctx, "123").Return(nil)

		err := ps.DeletePost(ctx, "user1", "123")

		assert.NoError(t, err)
		pr.AssertExpectations(t)
	})
}
