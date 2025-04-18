package services

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestHealthService_Check(t *testing.T) {
	t.Run("should return status ok when database is up", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPing().WillReturnError(nil)

		service := NewHealthService(db)

		status := service.Check(context.Background())

		assert.Equal(t, "ok", status.Status)
		assert.Equal(t, "ok", status.Dependencies["database"])
		assert.Empty(t, status.Error)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return status unhealthy when ping fails", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPing().WillReturnError(errors.New("db down"))

		service := NewHealthService(db)

		status := service.Check(context.Background())

		assert.Equal(t, "unhealthy", status.Status)
		assert.Equal(t, "down", status.Dependencies["database"])
		assert.Contains(t, status.Error, "db down")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
