package utils

import (
	"context"
	"database/sql"
	"errors"
)

type QueryExecutor interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func QueryRowScan[T any](ctx context.Context, db QueryExecutor, query string, scan func(*sql.Row) (*T, error), args ...any) (*T, error) {
	row := db.QueryRowContext(ctx, query, args...)

	result, err := scan(row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}
