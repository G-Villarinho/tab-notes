package handlers

import (
	"net/http"

	"github.com/g-villarinho/tab-notes-api/configs"
)

type EnvironmentHandler interface {
	GetEnvs(w http.ResponseWriter, r *http.Request)
}

type environmentHandler struct{}

func NewEnvironmentHandler() EnvironmentHandler {
	return &environmentHandler{}
}

func (e *environmentHandler) GetEnvs(w http.ResponseWriter, r *http.Request) {
	JSON(w, http.StatusOK, configs.Env)
}
