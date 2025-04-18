package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/g-villarinho/tab-notes-api/services"
)

type RegisterHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

type registerHandler struct {
	rs services.RegisterService
}

func NewRegisterHandler(rs services.RegisterService) RegisterHandler {
	return &registerHandler{
		rs: rs,
	}
}

func (rh *registerHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logger := slog.With(
		slog.String("handler", "register"),
		slog.String("method", "RegisterUser"),
	)

	var payload models.RegisterPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logger.Error("decode request body", "error", err)
		NoContent(w, http.StatusBadRequest)
		return
	}

	if err := rh.rs.RegisterUser(r.Context(), payload.Name, payload.Username, payload.Email); err != nil {
		if err == models.ErrEmailAlreadyExists {
			logger.Error("user already exists", "error", err)
			ErrorJSON(w, http.StatusConflict, "Já existe uma conta com esse e-mail.")
			return
		}

		if err == models.ErrUsernameAlreadyExists {
			logger.Error("user already exists", "error", err)
			ErrorJSON(w, http.StatusConflict, "Este nome de usuário já está em uso.")
			return
		}

		logger.Error("register user", "error", err)
		NoContent(w, http.StatusInternalServerError)
		return
	}

	NoContent(w, http.StatusCreated)
}
