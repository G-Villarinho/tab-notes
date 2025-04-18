package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterHandler_RegisterUser(t *testing.T) {
	t.Run("should return 400 if payload is invalid", func(t *testing.T) {
		rs := new(mocks.RegisterServiceMock)
		h := NewRegisterHandler(rs)

		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid json")))
		rr := httptest.NewRecorder()

		h.RegisterUser(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 409 if email already exists", func(t *testing.T) {
		rs := new(mocks.RegisterServiceMock)
		h := NewRegisterHandler(rs)

		rs.On("RegisterUser", mock.Anything, "João", "joaozinha-30", "joao@example.com").
			Return(models.ErrEmailAlreadyExists)

		payload := toJSON(t, models.RegisterPayload{
			Name:     "João",
			Username: "joaozinha-30",
			Email:    "joao@example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/register", payload)
		rr := httptest.NewRecorder()

		h.RegisterUser(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		rs.AssertExpectations(t)
	})

	t.Run("should return 409 if username already exists", func(t *testing.T) {
		rs := new(mocks.RegisterServiceMock)
		h := NewRegisterHandler(rs)

		rs.On("RegisterUser", mock.Anything, "João", "joaozinha-30", "joao@example.com").
			Return(models.ErrUsernameAlreadyExists)

		payload := toJSON(t, models.RegisterPayload{
			Name:     "João",
			Username: "joaozinha-30",
			Email:    "joao@example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/register", payload)
		rr := httptest.NewRecorder()

		h.RegisterUser(rr, req)

		assert.Equal(t, http.StatusConflict, rr.Code)
		rs.AssertExpectations(t)
	})

	t.Run("should return 500 on unexpected error", func(t *testing.T) {
		rs := new(mocks.RegisterServiceMock)
		h := NewRegisterHandler(rs)

		rs.On("RegisterUser", mock.Anything, "João", "joaozinha-30", "joao@example.com").
			Return(errors.New("unexpected error"))

		payload := toJSON(t, models.RegisterPayload{
			Name:     "João",
			Username: "joaozinha-30",
			Email:    "joao@example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/register", payload)
		rr := httptest.NewRecorder()

		h.RegisterUser(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		rs.AssertExpectations(t)
	})

	t.Run("should return 201 on success", func(t *testing.T) {
		rs := new(mocks.RegisterServiceMock)
		h := NewRegisterHandler(rs)

		rs.On("RegisterUser", mock.Anything, "João", "joaozinha-30", "joao@example.com").
			Return(nil)

		payload := toJSON(t, models.RegisterPayload{
			Name:     "João",
			Username: "joaozinha-30",
			Email:    "joao@example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/register", payload)
		rr := httptest.NewRecorder()

		h.RegisterUser(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		rs.AssertExpectations(t)
	})
}
