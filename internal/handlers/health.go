package handlers

import (
	"fmt"
	"net/http"
)

// HealthHandler — простой эндпоинт для проверки
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
