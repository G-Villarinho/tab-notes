package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/pkgs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthHandler_SendAuthenticationLink(t *testing.T) {
	rc := pkgs.NewRequestContext()

	t.Run("should return 400 if body is invalid", func(t *testing.T) {
		as := new(mocks.AuthServiceMock)
		h := NewAuthHandler(as, rc)

		req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewBufferString("invalid-json"))
		rr := httptest.NewRecorder()

		h.SendAuthenticationLink(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 200 even if user not found (silent)", func(t *testing.T) {
		as := new(mocks.AuthServiceMock)
		h := NewAuthHandler(as, rc)

		payload := models.SendAuthenticationLinkPayload{Email: "naoexiste@email.com"}
		body, _ := json.Marshal(payload)

		as.On("SendAuthenticationLink", mock.Anything, payload.Email).
			Return(models.ErrUserNotFound)

		req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		h.SendAuthenticationLink(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		as.AssertExpectations(t)
	})

	t.Run("should return 500 on error", func(t *testing.T) {
		as := new(mocks.AuthServiceMock)
		h := NewAuthHandler(as, rc)

		payload := models.SendAuthenticationLinkPayload{Email: "error@email.com"}
		body, _ := json.Marshal(payload)

		as.On("SendAuthenticationLink", mock.Anything, payload.Email).
			Return(assert.AnError)

		req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		h.SendAuthenticationLink(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		as.AssertExpectations(t)
	})

	t.Run("should return 200 on success", func(t *testing.T) {
		as := new(mocks.AuthServiceMock)
		h := NewAuthHandler(as, rc)

		payload := models.SendAuthenticationLinkPayload{Email: "ok@email.com"}
		body, _ := json.Marshal(payload)

		as.On("SendAuthenticationLink", mock.Anything, payload.Email).
			Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/authenticate", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		h.SendAuthenticationLink(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		as.AssertExpectations(t)
	})
}

func TestAuthHandler_AuthenticateFromLink(t *testing.T) {
	rc := pkgs.NewRequestContext()
	as := new(mocks.AuthServiceMock)
	handler := NewAuthHandler(as, rc)

	configs.Env.RedirectURL = "http://localhost:5173/"

	t.Run("should return 400 if token is missing", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/magic-link/authenticate", nil)
		rr := httptest.NewRecorder()

		handler.AuthenticateFromLink(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should redirect to fail if session not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/magic-link/authenticate?token=abc", nil)
		rr := httptest.NewRecorder()

		as.On("AuthenticateFromLink", mock.Anything, "abc").
			Return(nil, models.ErrSessionNotFound)

		handler.AuthenticateFromLink(rr, req)

		assert.Equal(t, http.StatusFound, rr.Code)
		assert.Equal(t, "http://localhost:5173/auth/fail?error=invalid_token", rr.Header().Get("Location"))
	})

	t.Run("should redirect to fail if session expired", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/magic-link/authenticate?token=expired", nil)
		rr := httptest.NewRecorder()

		as.On("AuthenticateFromLink", mock.Anything, "expired").
			Return(nil, models.ErrSessionExpired)

		handler.AuthenticateFromLink(rr, req)

		assert.Equal(t, http.StatusFound, rr.Code)
		assert.Equal(t, "http://localhost:5173/auth/fail?error=expired_token", rr.Header().Get("Location"))
	})

	t.Run("should return 500 if unknown error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/magic-link/authenticate?token=err", nil)
		rr := httptest.NewRecorder()

		as.On("AuthenticateFromLink", mock.Anything, "err").
			Return(nil, assert.AnError)

		handler.AuthenticateFromLink(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should set cookie and redirect to app", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/magic-link/authenticate?token=valid", nil)
		rr := httptest.NewRecorder()

		as.On("AuthenticateFromLink", mock.Anything, "valid").
			Return(&models.AuthResponse{Token: "abc.def.ghi"}, nil)

		handler.AuthenticateFromLink(rr, req)

		assert.Equal(t, http.StatusFound, rr.Code)
		assert.Equal(t, "http://localhost:5173/", rr.Header().Get("Location"))

		cookies := rr.Result().Cookies()
		var found bool
		for _, c := range cookies {
			if c.Name == "tabnews_id" && c.Value == "abc.def.ghi" {
				found = true
			}
		}
		assert.True(t, found, "Token cookie not set")
	})
}

func TestAuthHandler_Logout(t *testing.T) {
	t.Run("should return 401 if session ID is missing", func(t *testing.T) {
		rc := pkgs.NewRequestContext()
		as := new(mocks.AuthServiceMock)
		handler := NewAuthHandler(as, rc)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		rr := httptest.NewRecorder()

		handler.Logout(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return 200 OK if session not found", func(t *testing.T) {
		rc := pkgs.NewRequestContext()
		as := new(mocks.AuthServiceMock)
		handler := NewAuthHandler(as, rc)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		ctx := rc.SetSessionID(req.Context(), "sess-123")
		req = req.WithContext(ctx)

		as.On("Logout", mock.Anything, "sess-123").
			Return(models.ErrSessionNotFound)

		rr := httptest.NewRecorder()
		handler.Logout(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("should return 500 if logout returns error", func(t *testing.T) {
		rc := pkgs.NewRequestContext()
		as := new(mocks.AuthServiceMock)
		handler := NewAuthHandler(as, rc)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		ctx := rc.SetSessionID(req.Context(), "sess-123")
		req = req.WithContext(ctx)

		as.On("Logout", mock.Anything, "sess-123").
			Return(assert.AnError)

		rr := httptest.NewRecorder()
		handler.Logout(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should return 200 and clear cookie on success", func(t *testing.T) {
		rc := pkgs.NewRequestContext()
		as := new(mocks.AuthServiceMock)
		handler := NewAuthHandler(as, rc)

		req := httptest.NewRequest(http.MethodPost, "/logout", nil)
		ctx := rc.SetSessionID(req.Context(), "sess-123")
		req = req.WithContext(ctx)

		as.On("Logout", mock.Anything, "sess-123").
			Return(nil)

		rr := httptest.NewRecorder()
		handler.Logout(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		assert.Condition(t, func() bool {
			for _, c := range rr.Result().Cookies() {
				if c.Name == "tabnews_id" && c.MaxAge == -1 {
					return true
				}
			}
			return false
		}, "Expected 'tabnews_id' cookie to be deleted (MaxAge = -1)")
	})
}
