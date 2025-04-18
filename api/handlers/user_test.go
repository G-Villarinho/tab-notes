package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func toJSON(t *testing.T, v any) *bytes.Buffer {
	t.Helper()
	b, err := json.Marshal(v)
	assert.NoError(t, err)
	return bytes.NewBuffer(b)
}

func TestUserHandler_GetProfile(t *testing.T) {
	t.Run("should return 401 if userID not found in context", func(t *testing.T) {
		us := new(mocks.UserServiceMock)
		rc := new(mocks.RequestContextMock)
		rc.On("GetUserID", mock.Anything).Return("", false)

		h := NewUserHandler(rc, us)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		rr := httptest.NewRecorder()

		h.GetProfile(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		us.AssertExpectations(t)
	})

	t.Run("should return 404 if user not found", func(t *testing.T) {
		us := new(mocks.UserServiceMock)
		rc := new(mocks.RequestContextMock)

		rc.On("GetUserID", mock.Anything).Return("user-123", true)
		us.On("GetProfile", mock.Anything, "user-123").Return(nil, models.ErrUserNotFound)

		h := NewUserHandler(rc, us)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		rr := httptest.NewRecorder()

		h.GetProfile(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		us.AssertExpectations(t)
	})

	t.Run("should return 500 on internal error", func(t *testing.T) {
		us := new(mocks.UserServiceMock)
		rc := new(mocks.RequestContextMock)

		rc.On("GetUserID", mock.Anything).Return("user-123", true)
		us.On("GetProfile", mock.Anything, "user-123").Return(nil, errors.New("db error"))

		h := NewUserHandler(rc, us)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		rr := httptest.NewRecorder()

		h.GetProfile(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		us.AssertExpectations(t)
	})

	t.Run("should return 200 with user response", func(t *testing.T) {
		us := new(mocks.UserServiceMock)
		rc := new(mocks.RequestContextMock)

		rc.On("GetUserID", mock.Anything).Return("user-123", true)

		expected := &models.UserResponse{
			Name:      "Gabriel",
			Email:     "gabriel@example.com",
			Followers: 10,
			Following: 5,
		}

		us.On("GetProfile", mock.Anything, "user-123").Return(expected, nil)

		h := NewUserHandler(rc, us)

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		rr := httptest.NewRecorder()

		h.GetProfile(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Contains(t, rr.Body.String(), "gabriel@example.com")
		us.AssertExpectations(t)
	})
}
