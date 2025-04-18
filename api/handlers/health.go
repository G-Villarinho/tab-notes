package handlers

import (
	"net/http"

	"github.com/g-villarinho/tab-notes-api/services"
)

type HealthHandler struct {
	service services.HealthService
}

func NewHealthHandler(service services.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	status := h.service.Check(r.Context())

	code := http.StatusOK
	if status.Status != "ok" {
		code = http.StatusInternalServerError
	}

	JSON(w, code, status)
}
