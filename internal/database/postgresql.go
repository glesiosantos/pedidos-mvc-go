package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(ctx context.Context, url string) (*sql.DB, error) {
	url = strings.TrimSpace(url)
	if url == "" {
		return nil, errors.New("DATABASE_URL não definida")
	}

	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("erro ao configurar PostgreSQL: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao conectar ao PostgreSQL: %w", err)
	}

	return db, nil
}
