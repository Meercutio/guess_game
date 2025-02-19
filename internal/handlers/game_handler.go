package handlers

import (
	"fmt"
	"net/http"
)

// StartGameHandler — простой пример
func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Game started!")
}
