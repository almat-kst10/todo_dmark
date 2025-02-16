package db

import (
	"context"
	"database/sql"
	"fmt"
	"todo_dmark/configs"

	_ "github.com/lib/pq"
)

func NewDatabase(ctx context.Context, cfg *configs.Config) (*sql.DB, error) {
	d := `
		host=%s 
		port=%s 
		user=%s 
		dbname=%s 
		password=%s 
		sslmode=%s
		client_encoding=%s
	`
	dsn := fmt.Sprintf(d, cfg.DBHost, cfg.DBPort, cfg.DBDriver, cfg.DBName, cfg.DBPsw, cfg.DBSslmode, cfg.DBEncoding)
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
