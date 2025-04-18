package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type PostHandler interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
	UnlikePost(w http.ResponseWriter, r *http.Request)
	GetPostsByUsername(w http.ResponseWriter, r *http.Request)
	GetPostsByAuthorID(w http.ResponseWriter, r *http.Request)
}

type postHandler struct {
	rc pkgs.RequestContext
	ps services.PostService
}

func NewPostHandler(
	requestContext pkgs.RequestContext,
	postService services.PostService) PostHandler {
	return &postHandler{
		rc: requestContext,
		ps: postService,
	}
}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "CreatePost"),
	)

	var payload models.CreatePostPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("decode payload", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	response, err := p.ps.CreatePost(r.Context(), userID, payload.Title, payload.Content)
	if err != nil {
		logger.Error("create post", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusCreated, response)
}

func (p *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "GetPostByID"),
	)

	postID := r.PathValue("postId")
	if postID == "" {
		logger.Error("get post by id", "error", "post id not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	post, err := p.ps.GetPostByID(r.Context(), userID, postID)
	if err != nil {
		if err == models.ErrPostNotFound {
			logger.Error("get post by id", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("get post by id", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, post)
}

func (p *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "DeletePost"),
	)

	postID := r.PathValue("postId")
	if postID == "" {
		logger.Error("delete post", "error", "post id not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := p.ps.DeletePost(r.Context(), userID, postID); err != nil {
		if err == models.ErrPostNotFound {
			logger.Error("delete post", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		if err == models.ErrPostNotBelongToUser {
			logger.Error("delete post", "error", err)
			NoContent(w, http.StatusForbidden)
			return
		}

		logger.Error("delete post", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}

func (p *postHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "UpdatePost"),
	)

	postID := r.PathValue("postId")
	if postID == "" {
		logger.Error("update post", "error", "post id not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	var payload models.UpdatePostPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("decode payload", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	if err := p.ps.UpdatePost(r.Context(), userID, postID, payload.Title, payload.Content); err != nil {
		if err == models.ErrPostNotFound {
			logger.Error("update post", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		if err == models.ErrPostNotBelongToUser {
			logger.Error("update post", "error", err)
			NoContent(w, http.StatusForbidden)
			return
		}

		logger.Error("update post", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}

func (p *postHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "LikePost"),
	)

	postID := r.PathValue("postId")
	if postID == "" {
		logger.Error("like post", "error", "post id not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := p.ps.LikePost(r.Context(), userID, postID); err != nil {
		if err == models.ErrPostNotFound {
			logger.Error("like post", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("like post", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}

func (p *postHandler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "UnlikePost"),
	)

	postID := r.PathValue("postId")
	if postID == "" {
		logger.Error("unlike post", "error", "post id not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := p.ps.UnlikePost(r.Context(), userID, postID); err != nil {
		if err == models.ErrPostNotFound {
			logger.Error("unlike post", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("unlike post", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}

func (p *postHandler) GetPostsByUsername(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "GetPostsByUsername"),
	)

	username := r.PathValue("username")
	if username == "" {
		logger.Error("get posts by username", "error", "username not found in query params")
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	posts, err := p.ps.GetPostsByUsername(r.Context(), userID, username)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Error("get posts by username", "error", err)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("get posts by username", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(posts) == 0 {
		logger.Info("get posts by username", "info", "no posts found")
	}

	JSON(w, http.StatusOK, posts)
}

func (p *postHandler) GetPostsByAuthorID(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "post"),
		slog.String("method", "GetPostsByAuthorID"),
	)

	authorID, ok := p.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("get user id from context", "error", "user id not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	posts, err := p.ps.GetPostsByAuthorID(r.Context(), authorID)
	if err != nil {
		logger.Error("get posts by author id", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(posts) == 0 {
		logger.Info("get posts by author id", "info", "no posts found")
	}

	JSON(w, http.StatusOK, posts)
}
