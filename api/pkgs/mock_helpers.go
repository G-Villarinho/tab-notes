package pkgs

import (
	"database/sql"
	"time"

	"github.com/g-villarinho/tab-notes-api/models"
	"github.com/stretchr/testify/mock"
)

func MockAnyTime() any {
	return mock.MatchedBy(func(t time.Time) bool { return !t.IsZero() })
}

func SQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func SQLNullTimeZero() sql.NullTime {
	return sql.NullTime{Valid: false}
}

func MockSessionWithToken(token string) any {
	return mock.MatchedBy(func(s *models.Session) bool {
		return s != nil && s.Token == token
	})
}
