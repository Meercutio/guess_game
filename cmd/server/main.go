package main

import (
	"fmt"
	"guess_game/internal/handlers"
	"log"
	"net/http"
)

// healthHandler — простой обработчик, возвращающий "OK"
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
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
