package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/hermes-mailer/models"
	"github.com/hermes-mailer/services"
)

type EmailHandler interface {
	SendEmail(w http.ResponseWriter, r *http.Request)
}

type emailHandler struct {
	es services.EmailService
}

func NewEmailHandler(es services.EmailService) EmailHandler {
	return &emailHandler{
		es: es,
	}
}

func (e *emailHandler) SendEmail(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "email"),
		slog.String("method", "SendEmail"),
	)

	var payload models.Email
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("decode request body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := e.es.Enqueue(r.Context(), payload); err != nil {
		logger.Error("enqueue email", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
