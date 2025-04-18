package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g-villarinho/tab-notes-api/mocks"
	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHealthHandler_Check(t *testing.T) {
	t.Run("should return 200 when status is ok", func(t *testing.T) {
		service := new(mocks.HealthServiceMock)
		handler := NewHealthHandler(service)

		service.On("Check", mock.Anything).
			Return(models.HealthStatus{
				Status:       "ok",
				Dependencies: map[string]string{"database": "ok"},
			})

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		handler.Check(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		service.AssertExpectations(t)
	})

	t.Run("should return 500 when status is unhealthy", func(t *testing.T) {
		service := new(mocks.HealthServiceMock)
		handler := NewHealthHandler(service)

		service.On("Check", mock.Anything).
			Return(models.HealthStatus{
				Status:       "unhealthy",
				Dependencies: map[string]string{"database": "down"},
				Error:        "db unreachable",
			})

		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		rr := httptest.NewRecorder()

		handler.Check(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		service.AssertExpectations(t)
	})
}
