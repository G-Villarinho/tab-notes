package services

import (
	"context"
	"errors"
	"testing"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateLike(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		like := &models.Like{
			UserID: "user-123",
			PostID: "post-456",
		}

		lr.
			On("CreateLike", ctx, like).
			Return(errors.New("insert failed"))

		err := ls.LikePost(ctx, "user-123", "post-456")

		assert.ErrorContains(t, err, "error creating like")
		lr.AssertExpectations(t)
	})

	t.Run("should like post successfully", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		like := &models.Like{
			UserID: "user-123",
			PostID: "post-456",
		}

		lr.
			On("CreateLike", ctx, like).
			Return(nil)

		err := ls.LikePost(ctx, "user-123", "post-456")

		assert.NoError(t, err)
		lr.AssertExpectations(t)
	})
}

func TestDeleteLike(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails to delete like", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		like := &models.Like{
			UserID: "user-123",
			PostID: "post-456",
		}

		lr.
			On("DeleteLike", ctx, like).
			Return(errors.New("delete error"))

		err := ls.UnlikePost(ctx, "user-123", "post-456")

		assert.ErrorContains(t, err, "error deleting like")
		lr.AssertExpectations(t)
	})

	t.Run("should unlike post successfully", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		like := &models.Like{
			UserID: "user-123",
			PostID: "post-456",
		}

		lr.
			On("DeleteLike", ctx, like).
			Return(nil)

		err := ls.UnlikePost(ctx, "user-123", "post-456")

		assert.NoError(t, err)
		lr.AssertExpectations(t)
	})
}

func TestCheckLikes(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		postIDs := []string{"p1", "p2"}

		lr.
			On("GetLikedPostIDs", ctx, "user-123", postIDs).
			Return(nil, errors.New("db error"))

		likedMap, err := ls.CheckLikes(ctx, "user-123", postIDs)

		assert.Nil(t, likedMap)
		assert.ErrorContains(t, err, "db error")
		lr.AssertExpectations(t)
	})

	t.Run("should return correct liked map", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		postIDs := []string{"p1", "p2", "p3"}
		liked := []string{"p1", "p3"}

		lr.
			On("GetLikedPostIDs", ctx, "user-123", postIDs).
			Return(liked, nil)

		likedMap, err := ls.CheckLikes(ctx, "user-123", postIDs)

		assert.NoError(t, err)
		assert.Len(t, likedMap, 2)
		assert.True(t, likedMap["p1"])
		assert.True(t, likedMap["p3"])
		assert.False(t, likedMap["p2"]) // implícito: não está presente
		lr.AssertExpectations(t)
	})
}

func TestCheckLike(t *testing.T) {
	ctx := context.Background()

	t.Run("should return error if repository fails", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		lr.
			On("CheckLike", ctx, "user-123", "post-456").
			Return(false, errors.New("db error"))

		liked, err := ls.CheckLike(ctx, "user-123", "post-456")

		assert.False(t, liked)
		assert.ErrorContains(t, err, "check like")
		lr.AssertExpectations(t)
	})

	t.Run("should return true when user liked the post", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		lr.
			On("CheckLike", ctx, "user-123", "post-456").
			Return(true, nil)

		liked, err := ls.CheckLike(ctx, "user-123", "post-456")

		assert.NoError(t, err)
		assert.True(t, liked)
		lr.AssertExpectations(t)
	})

	t.Run("should return false when user has not liked the post", func(t *testing.T) {
		lr := new(mocks.LikeRepositoryMock)
		ls := NewLikeService(lr)

		lr.
			On("CheckLike", ctx, "user-123", "post-456").
			Return(false, nil)

		liked, err := ls.CheckLike(ctx, "user-123", "post-456")

		assert.NoError(t, err)
		assert.False(t, liked)
		lr.AssertExpectations(t)
	})
}
