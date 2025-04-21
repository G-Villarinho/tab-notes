package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type SessionHandler interface {
	GetUserSessions(w http.ResponseWriter, r *http.Request)
	RevokeSession(w http.ResponseWriter, r *http.Request)
	RevokeAllSessions(w http.ResponseWriter, r *http.Request)
}

type sessionHandler struct {
	rc pkgs.RequestContext
	ss services.SessionService
}

func NewSessionHandler(requestContext pkgs.RequestContext, sessionService services.SessionService) SessionHandler {
	return &sessionHandler{
		rc: requestContext,
		ss: sessionService,
	}
}

func (s *sessionHandler) GetUserSessions(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "user"),
		slog.String("method", "GetUserSessions"),
	)

	userID, ok := s.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	currentSessionID, ok := s.rc.GetSessionID(r.Context())
	if !ok {
		logger.Error("sessionID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	response, err := s.ss.GetUserSessions(r.Context(), userID, currentSessionID)
	if err != nil {
		if err == models.ErrUserNotFound {
			logger.Error("user not found", "userID", userID)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error getting user sessions", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	if len(response) == 0 {
		logger.Info("no sessions found for user", "userID", userID)
	}

	JSON(w, http.StatusOK, response)
}

func (s *sessionHandler) RevokeSession(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "session"),
		slog.String("method", "RevokeSession"),
	)

	userID, ok := s.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		DeleteTokenCookie(w)
		NoContent(w, http.StatusUnauthorized)
		return
	}

	sessionID := r.PathValue("sessionId")
	if sessionID == "" {
		logger.Error("session_id not found in query")
		NoContent(w, http.StatusBadRequest)
		return
	}

	if err := s.ss.RevokeUserSession(r.Context(), userID, sessionID); err != nil {
		if err == models.ErrSessionNotFound {
			logger.Error("session not found", "sessionID", sessionID)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error revoking session", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}

func (s *sessionHandler) RevokeAllSessions(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "session"),
		slog.String("method", "RevokeAllSessions"),
	)

	var payload models.RevokeAllSessionsPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("error decoding request body", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	userID, ok := s.rc.GetUserID(r.Context())
	if !ok {
		logger.Error("userID not found in context")
		NoContent(w, http.StatusUnauthorized)
		DeleteTokenCookie(w)
		return
	}

	currentSessionID, ok := s.rc.GetSessionID(r.Context())
	if !ok {
		logger.Error("sessionID not found in context")
		NoContent(w, http.StatusUnauthorized)
		DeleteTokenCookie(w)
		return
	}

	err := s.ss.RevokeAllUserSessions(r.Context(), userID, currentSessionID, payload.RevokeCurrent)
	if err != nil {
		if err == models.ErrSessionNotFound {
			logger.Error("session not found", "sessionID", userID)
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("error revoking all sessions", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusNoContent)
}
