package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // драйвер pgx
)

// InitPostgres — инициализируем соединение к PostgreSQL
func InitPostgres(host, port, user, pass, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	log.Printf("Connected to PostgreSQL (%s:%s/%s)\n", host, port, dbname)
	return db, nil
}

// TestDBConnection — пример тестового запроса (необязательно)
func TestDBConnection(ctx context.Context, db *sql.DB) error {
	var version string
	err := db.QueryRowContext(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		return err
	}
	log.Println("PostgreSQL version:", version)
	return nil
}
