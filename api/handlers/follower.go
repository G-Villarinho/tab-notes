package handlers

import (
	"log/slog"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type FollowerHandler interface {
	FollowUser(w http.ResponseWriter, r *http.Request)
	UnfollowUser(w http.ResponseWriter, r *http.Request)
	GetFollowers(w http.ResponseWriter, r *http.Request)
	GetFollowing(w http.ResponseWriter, r *http.Request)
	GetMyFollowers(w http.ResponseWriter, r *http.Request)
	GetMyFollowing(w http.ResponseWriter, r *http.Request)
}

type followerHandler struct {
	rc pkgs.RequestContext
	fs services.FollowerService
}

func NewFollowerHandler(rc pkgs.RequestContext, fs services.FollowerService) FollowerHandler {
	return &followerHandler{
		rc: rc,
		fs: fs,
	}
}

func (f *followerHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "FollowUser"),
	)

	username := r.PathValue("username")

	if username == "" {
		NoContent(w, http.StatusBadRequest)
		return
	}

	followerID, ok := f.rc.GetUserID(r.Context())
	if !ok {
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := f.fs.FollowUser(r.Context(), followerID, username); err != nil {
		switch err {
		case models.ErrUserNotFound:
			logger.Warn("user not found")
			NoContent(w, http.StatusNotFound)
			return
		case models.ErrCannotFollowSelf:
			logger.Warn("cannot follow self")
			NoContent(w, http.StatusForbidden)
			return
		default:
			logger.Error("error following user", slog.String("error", err.Error()))
			NoContent(w, http.StatusInternalServerError)
			return
		}
	}

}

func (f *followerHandler) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "UnfollowUser"),
	)

	username := r.PathValue("username")

	if username == "" {
		NoContent(w, http.StatusBadRequest)
		return
	}

	followerID, ok := f.rc.GetUserID(r.Context())
	if !ok {
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := f.fs.UnfollowUser(r.Context(), followerID, username); err != nil {
		switch err {
		case models.ErrUserNotFound:
			logger.Warn("user not found")
			NoContent(w, http.StatusNotFound)
			return
		case models.ErrCannotUnfollowSelf:
			logger.Warn("cannot unfollow self")
			NoContent(w, http.StatusForbidden)
			return
		default:
			logger.Error("error unfollowing user", slog.String("error", err.Error()))
			NoContent(w, http.StatusInternalServerError)
			return
		}
	}

	NoContent(w, http.StatusNoContent)
}

func (f *followerHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "GetFollowers"),
	)

	username := r.PathValue("username")

	if username == "" {
		NoContent(w, http.StatusBadRequest)
		return
	}

	followers, err := f.fs.GetFollowers(r.Context(), username)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Warn("user not found")
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error getting followers", slog.String("error", err.Error()))
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(followers) == 0 {
		logger.Info("no followers found")
	}

	JSON(w, http.StatusOK, followers)
}

func (f *followerHandler) GetFollowing(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "GetFollowing"),
	)

	username := r.PathValue("username")

	if username == "" {
		NoContent(w, http.StatusBadRequest)
		return
	}

	following, err := f.fs.GetFollowing(r.Context(), username)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Warn("user not found")
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error getting following", slog.String("error", err.Error()))
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(following) == 0 {
		logger.Info("no following found")
	}

	JSON(w, http.StatusOK, following)
}

func (f *followerHandler) GetMyFollowers(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "GetMyFollowers"),
	)

	userID, ok := f.rc.GetUserID(r.Context())
	if !ok {
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	followers, err := f.fs.GetMyFollowers(r.Context(), userID)
	if err != nil {
		logger.Error("error getting followers", slog.String("error", err.Error()))
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(followers) == 0 {
		logger.Info("no followers found")
	}

	JSON(w, http.StatusOK, followers)
}

func (f *followerHandler) GetMyFollowing(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "follower"),
		slog.String("method", "GetMyFollowing"),
	)

	userID, ok := f.rc.GetUserID(r.Context())
	if !ok {
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	following, err := f.fs.GetMyFollowing(r.Context(), userID)
	if err != nil {
		logger.Error("error getting following", slog.String("error", err.Error()))
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(following) == 0 {
		logger.Info("no following found")
	}

	JSON(w, http.StatusOK, following)
}
