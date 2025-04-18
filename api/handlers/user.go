package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type UserHandler interface {
	GetProfile(w http.ResponseWriter, r *http.Request)
	SearchUsers(w http.ResponseWriter, r *http.Request)
	GetProfileByUsername(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	rc pkgs.RequestContext
	us services.UserService
}

func NewUserHandler(requestContext pkgs.RequestContext, userService services.UserService) UserHandler {
	return &userHandler{
		us: userService,
		rc: requestContext,
	}
}

func (u *userHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "user"),
		slog.String("method", "GetProfile"),
	)

	userID, ok := u.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		NoContent(w, http.StatusUnauthorized)
		DeleteTokenCookie(w)
		return
	}

	response, err := u.us.GetProfile(r.Context(), userID)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Error("user not found", "userID", userID)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error getting user profile", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, response)
}

func (u *userHandler) SearchUsers(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "user"),
		slog.String("method", "SearchUsers"),
	)

	query := strings.ToLower(r.URL.Query().Get("q"))
	if query == "" {
		logger.Error("query not found in query")
		NoContent(w, http.StatusBadRequest)
		return
	}

	users, err := u.us.SearchUsers(r.Context(), query)
	if err != nil {
		logger.Error("error searching users", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		logger.Info("no users found")
	}

	JSON(w, http.StatusOK, users)
}

func (u *userHandler) GetProfileByUsername(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "user"),
		slog.String("method", "GetProfileByUsername"),
	)

	username := strings.ToLower(r.PathValue("username"))
	if username == "" {
		logger.Error("username not found in query")
		NoContent(w, http.StatusBadRequest)
		return
	}

	viewerID, ok := u.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("viewerID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	response, err := u.us.GetProfileByUsername(r.Context(), username, viewerID)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Error("user not found", "username", username)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error getting user profile by username", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	JSON(w, http.StatusOK, response)
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "user"),
		slog.String("method", "UpdateUser"),
	)

	var payload models.UpdateUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("error decoding request body", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := u.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := u.us.UpdateUser(r.Context(), userID, payload.Name, payload.Username); err != nil {
		if err == models.ErrUserNotFound {
			logger.Error("user not found", "userID", userID)
			NoContent(w, http.StatusNotFound)
			return
		}

		if err == models.ErrUsernameAlreadyExists {
			logger.Error("username already exists", "username", payload.Username)
			NoContent(w, http.StatusConflict)
			return
		}

		logger.Error("error updating user", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusOK)
}
