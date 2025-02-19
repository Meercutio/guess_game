package main

import (
	"context"
	"fmt"
	"guess_game/internal/handlers"
	"guess_game/internal/storage"
	"log"
	"net/http"
	"os"
	"time"
)

// healthHandler — простой обработчик, возвращающий "OK"
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// Считываем env-переменные
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// Подключаемся к PostgreSQL
	db, err := storage.NewDB(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Conn.Close()

	// Тестовая выборка
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.ExampleQuery(ctx); err != nil {
		log.Println("Error in ExampleQuery:", err)
	}

	// Регистрируем обработчик
	http.HandleFunc("/health", healthHandler)

	// Пример эндпоинта для старта игры
	http.HandleFunc("/game/start", handlers.StartGameHandler)

	// Запускаем сервер на порту 8080
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
