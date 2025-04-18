package storages

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/g-villarinho/tab-notes-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(ctx context.Context) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC",
		configs.Env.DB.DBUser,
		configs.Env.DB.DBPassword,
		configs.Env.DB.DBHost,
		configs.Env.DB.DBPort,
		configs.Env.DB.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %v", err)
	}

	return db, nil
}
