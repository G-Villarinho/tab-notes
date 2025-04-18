package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type LikeService interface {
	LikePost(ctx context.Context, userID, postID string) error
	UnlikePost(ctx context.Context, userID, postID string) error
	CheckLikes(ctx context.Context, userID string, postIDs []string) (map[string]bool, error)
	CheckLike(ctx context.Context, userID, postID string) (bool, error)
}

type likeService struct {
	lr repositories.LikeRepository
}

func NewLikeService(likeRepository repositories.LikeRepository) LikeService {
	return &likeService{
		lr: likeRepository,
	}
}

func (l *likeService) LikePost(ctx context.Context, userID string, postID string) error {
	like := &models.Like{
		UserID: userID,
		PostID: postID,
	}

	if err := l.lr.CreateLike(ctx, like); err != nil {
		return fmt.Errorf("error creating like: %w", err)
	}

	return nil
}

func (l *likeService) UnlikePost(ctx context.Context, userID string, postID string) error {
	like := &models.Like{
		UserID: userID,
		PostID: postID,
	}

	if err := l.lr.DeleteLike(ctx, like); err != nil {
		return fmt.Errorf("error deleting like: %w", err)
	}

	return nil
}

func (l *likeService) CheckLikes(ctx context.Context, userID string, postIDs []string) (map[string]bool, error) {
	IDs, err := l.lr.GetLikedPostIDs(ctx, userID, postIDs)
	if err != nil {
		return nil, err
	}

	likedMap := make(map[string]bool, len(IDs))
	for _, id := range IDs {
		likedMap[id] = true
	}

	return likedMap, nil
}

func (l *likeService) CheckLike(ctx context.Context, userID, postID string) (bool, error) {
	liked, err := l.lr.CheckLike(ctx, userID, postID)
	if err != nil {
		return false, fmt.Errorf("check like: %w", err)
	}

	return liked, nil
}
