package services

import (
	"context"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type FeedService interface {
	GetFeed(ctx context.Context, userID string, limit, offset int) ([]*models.FeedPostResponse, error)
}

type feedService struct {
	ls LikeService
	fr repositories.FeedRepository
}

func NewFeedService(
	likeService LikeService,
	feedRepository repositories.FeedRepository) FeedService {
	return &feedService{
		ls: likeService,
		fr: feedRepository,
	}
}

func (f *feedService) GetFeed(ctx context.Context, userID string, limit int, offset int) ([]*models.FeedPostResponse, error) {
	feed, err := f.fr.GetFeed(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(feed) == 0 {
		return feed, nil
	}

	postIDs := make([]string, len(feed))
	for i, post := range feed {
		postIDs[i] = post.PostID
	}

	likedMap, err := f.ls.CheckLikes(ctx, userID, postIDs)
	if err != nil {
		return nil, err
	}

	for _, post := range feed {
		post.LikedByUser = likedMap[post.PostID]
	}

	return feed, nil
}
