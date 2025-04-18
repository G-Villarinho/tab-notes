package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
)

type HealthService interface {
	Check(ctx context.Context) models.HealthStatus
}

type healthService struct {
	db *sql.DB
}

func NewHealthService(db *sql.DB) HealthService {
	return &healthService{db: db}
}

func (s *healthService) Check(ctx context.Context) models.HealthStatus {
	status := models.HealthStatus{
		Status:       "ok",
		Dependencies: make(map[string]string),
	}

	dbCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := s.db.PingContext(dbCtx); err != nil {
		status.Status = "unhealthy"
		status.Dependencies["database"] = "down"
		status.Error = err.Error()
		return status
	}

	status.Dependencies["database"] = "ok"
	return status
}
