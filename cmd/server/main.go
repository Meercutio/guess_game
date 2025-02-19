package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"guess_game/internal/handlers"
	"guess_game/internal/storage"
)

func main() {
	ctx := context.Background()

	// 1. Считываем окружение для PostgreSQL
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASS", "postgres")
	dbName := getEnv("DB_NAME", "guessdb")

	// 2. Инициализация PostgreSQL
	db, err := storage.InitPostgres(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatalf("Failed to init Postgres: %v", err)
	}
	// Проверка соединения (необязательно)
	if err := storage.TestDBConnection(ctx, db); err != nil {
		log.Println("Postgres test query error:", err)
	}

	// 3. Инициализация Redis
	redisHost := getEnv("REDIS_HOST", "localhost:6379")
	redisPass := getEnv("REDIS_PASSWORD", "")
	redisClient, err := storage.InitRedis(redisHost, redisPass, 0)
	if err != nil {
		log.Fatalf("Failed to init Redis: %v", err)
	}
	// Проверка Redis (необязательно)
	if err := storage.TestRedisConnection(ctx, redisClient); err != nil {
		log.Println("Redis test error:", err)
	}

	// 4. Регистрируем эндпоинты с помощью стандартной библиотеки
	http.HandleFunc("/health", handlers.HealthHandler)

	// Для /game/start нам нужно замкнуть на db и redisClient:
	http.HandleFunc("/game/start", handlers.StartGameHandler(db, redisClient))

	// 5. Запускаем сервер на :8080
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.DefaultServeMux, // используем стандартный ServeMux
	}

	log.Println("Starting server on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// getEnv — вспомогательная функция для чтения переменных окружения с дефолтом
func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
