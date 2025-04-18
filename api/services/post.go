package services

import (
	"context"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/repositories"
)

type PostService interface {
	CreatePost(ctx context.Context, userID string, title string, content string) (*models.PostResponse, error)
	LikePost(ctx context.Context, userID string, postID string) error
	UnlikePost(ctx context.Context, userID string, postID string) error
	GetPostByID(ctx context.Context, userID string, ID string) (*models.PostResponse, error)
	DeletePost(ctx context.Context, userID string, ID string) error
	UpdatePost(ctx context.Context, userID string, ID string, title string, content string) error
	GetPostsByUsername(ctx context.Context, userID string, username string) ([]*models.PostResponse, error)
	GetPostsByAuthorID(ctx context.Context, authorID string) ([]*models.PostResponse, error)
}

type postService struct {
	ls LikeService
	pr repositories.PostRepository
	ur repositories.UserRepository
}

func NewPostService(
	likeService LikeService,
	postRepository repositories.PostRepository,
	userRepository repositories.UserRepository) PostService {
	return &postService{
		ls: likeService,
		pr: postRepository,
		ur: userRepository,
	}
}

func (p *postService) CreatePost(ctx context.Context, userID string, title string, content string) (*models.PostResponse, error) {
	post := &models.Post{
		Title:    title,
		Content:  content,
		AuthorID: userID,
	}

	if err := p.pr.CreatePost(ctx, post); err != nil {
		return nil, fmt.Errorf("create post: %w", err)
	}

	return &models.PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Likes:     post.Likes,
		CreatedAt: post.CreatedAt,
	}, nil
}

func (p *postService) LikePost(ctx context.Context, userID string, postID string) error {
	post, err := p.pr.GetPostByID(ctx, postID)
	if err != nil {
		return fmt.Errorf("get post by id: %w", err)
	}

	if post == nil {
		return models.ErrPostNotFound
	}

	if err := p.ls.LikePost(ctx, userID, postID); err != nil {
		return fmt.Errorf("like post: %w", err)
	}

	return nil
}

func (p *postService) UnlikePost(ctx context.Context, userID string, postID string) error {
	post, err := p.pr.GetPostByID(ctx, postID)
	if err != nil {
		return fmt.Errorf("get post by id: %w", err)
	}

	if post == nil {
		return models.ErrPostNotFound
	}

	if err := p.ls.UnlikePost(ctx, userID, postID); err != nil {
		return fmt.Errorf("unlike post: %w", err)
	}

	return nil
}

func (p *postService) GetPostByID(ctx context.Context, userID string, ID string) (*models.PostResponse, error) {
	post, err := p.pr.GetPostByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("get post by id: %w", err)
	}

	if post == nil {
		return nil, nil
	}

	likedByUser, err := p.ls.CheckLike(ctx, userID, post.ID)
	if err != nil {
		return nil, fmt.Errorf("check like: %w", err)
	}

	postResponse := &models.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Content:     post.Content,
		Likes:       post.Likes,
		LikedByUser: likedByUser,
		CreatedAt:   post.CreatedAt,
	}

	return postResponse, nil
}

func (p *postService) DeletePost(ctx context.Context, userID string, ID string) error {
	post, err := p.pr.GetPostByID(ctx, ID)
	if err != nil {
		return fmt.Errorf("get post by id: %w", err)
	}

	if post == nil {
		return models.ErrPostNotFound
	}

	if post.AuthorID != userID {
		return models.ErrPostNotBelongToUser
	}

	if err := p.pr.DeletePost(ctx, ID); err != nil {
		return fmt.Errorf("delete post: %w", err)
	}

	return nil
}

func (p *postService) UpdatePost(ctx context.Context, userID string, ID string, title string, content string) error {
	post, err := p.pr.GetPostByID(ctx, ID)
	if err != nil {
		return fmt.Errorf("get post by id %s: %w", ID, err)
	}

	if post == nil {
		return models.ErrPostNotFound
	}

	if post.AuthorID != userID {
		return models.ErrPostNotBelongToUser
	}

	post.Title = title
	post.Content = content

	if err := p.pr.UpdatePost(ctx, post); err != nil {
		return fmt.Errorf("update post %s: %w", ID, err)
	}

	return nil
}

func (p *postService) GetPostsByUsername(ctx context.Context, userID string, username string) ([]*models.PostResponse, error) {
	author, err := p.ur.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username %s: %w", username, err)
	}

	if author == nil {
		return nil, models.ErrUserNotFound
	}

	posts, err := p.pr.GetPostsByAuthorID(ctx, author.ID)
	if err != nil {
		return nil, fmt.Errorf("get posts by author id %s: %w", author.ID, err)
	}

	if len(posts) == 0 {
		return []*models.PostResponse{}, nil
	}

	postIDs := make([]string, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	likedMap, err := p.ls.CheckLikes(ctx, userID, postIDs)
	if err != nil {
		return nil, fmt.Errorf("check likes: %w", err)
	}

	postResponses := make([]*models.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = &models.PostResponse{
			ID:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Likes:       post.Likes,
			LikedByUser: likedMap[post.ID],
			CreatedAt:   post.CreatedAt,
		}
	}

	return postResponses, nil
}

func (p *postService) GetPostsByAuthorID(ctx context.Context, authorID string) ([]*models.PostResponse, error) {
	posts, err := p.pr.GetPostsByAuthorID(ctx, authorID)
	if err != nil {
		return nil, fmt.Errorf("get posts by author id %s: %w", authorID, err)
	}

	postIDs := make([]string, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	likedMap, err := p.ls.CheckLikes(ctx, authorID, postIDs)
	if err != nil {
		return nil, fmt.Errorf("check likes: %w", err)
	}

	postResponses := make([]*models.PostResponse, len(posts))
	for i, post := range posts {
		postResponses[i] = &models.PostResponse{
			ID:          post.ID,
			Title:       post.Title,
			Content:     post.Content,
			Likes:       post.Likes,
			LikedByUser: likedMap[post.ID],
			CreatedAt:   post.CreatedAt,
		}
	}

	return postResponses, nil
}
