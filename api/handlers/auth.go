package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/g-villarinho/tab-notes-api/services"
)

type AuthHandler interface {
	SendAuthenticationLink(w http.ResponseWriter, r *http.Request)
	AuthenticateFromLink(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	rc pkgs.RequestContext
	as services.AuthService
}

func NewAuthHandler(authService services.AuthService, requestContext pkgs.RequestContext) AuthHandler {
	return &authHandler{
		as: authService,
		rc: requestContext,
	}
}

func (a *authHandler) SendAuthenticationLink(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "auth"),
		slog.String("method", "SendAuthenticationLink"),
	)

	var sendAuthenticationLinkPayload models.SendAuthenticationLinkPayload
	if err := json.NewDecoder(r.Body).Decode(&sendAuthenticationLinkPayload); err != nil {
		logger.Error("failed to decode request body", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	if err := a.as.SendAuthenticationLink(r.Context(), sendAuthenticationLinkPayload.Email); err != nil {
		if err == models.ErrUserNotFound {
			logger.Warn("user not found")
			NoContent(w, http.StatusNotFound)
			return
		}

		logger.Error("send authentication link", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusOK)
}

func (a *authHandler) AuthenticateFromLink(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "auth"),
		slog.String("method", "AuthenticateFromLink"),
	)

	token := r.URL.Query().Get("token")
	if token == "" {
		logger.Error("missing token")
		NoContent(w, http.StatusBadRequest)
		return
	}

	authResponse, err := a.as.AuthenticateFromLink(r.Context(), token)
	if err != nil {
		if err == models.ErrSessionNotFound {
			logger.Warn("invalid token (session not found)")
			http.Redirect(w, r, configs.Env.RedirectURL+"auth/fail?error=invalid_token", http.StatusFound)
			return
		}

		if err == models.ErrSessionExpired {
			logger.Warn("expired token")
			http.Redirect(w, r, configs.Env.RedirectURL+"auth/fail?error=expired_token", http.StatusFound)
			return
		}

		logger.Error("authenticate from link", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	SetTokenCookie(w, authResponse.Token)
	http.Redirect(w, r, configs.Env.RedirectURL, http.StatusFound)
}

func (a *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "auth"),
		slog.String("method", "Logout"),
	)

	sessionID, ok := a.rc.GetSessionID(r.Context())
	if !ok {
		logger.Error("missing session ID")
		NoContent(w, http.StatusUnauthorized)
		return
	}

	if err := a.as.Logout(r.Context(), sessionID); err != nil {
		if err == models.ErrSessionNotFound {
			logger.Warn("session not found (silenced)")
			NoContent(w, http.StatusOK)
		}

		logger.Error("logout", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	DeleteTokenCookie(w)
	NoContent(w, http.StatusOK)
}
