package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // импорт драйвера
)

type DB struct {
	Conn *sql.DB
}

// NewDB инициализирует подключение к PostgreSQL
func NewDB(host, port, user, pass, dbname string) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, dbname)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Проверим соединение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &DB{Conn: db}, nil
}

// ExampleQuery демонстрирует простую выборку
func (db *DB) ExampleQuery(ctx context.Context) error {
	var version string
	err := db.Conn.QueryRowContext(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		return err
	}
	fmt.Println("PostgreSQL version:", version)
	return nil
}
