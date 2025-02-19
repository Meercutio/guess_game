package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// StartGameHandler возвращает функцию-обработчик, замкнутую на db и redisClient.
func StartGameHandler(db *sql.DB, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Допустим, создаём запись в таблице users (упростим — без параметров из запроса):
		username := "player123"
		_, err := db.ExecContext(r.Context(),
			"INSERT INTO users(username) VALUES($1)", username)
		if err != nil {
			log.Printf("DB insert error: %v\n", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		// Запишем состояние игры в Redis
		gameSessionID := "game-001"
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		err = rdb.Set(ctx, gameSessionID, "in_progress", 30*time.Minute).Err()
		if err != nil {
			log.Printf("Redis set error: %v\n", err)
			http.Error(w, "Failed to start game", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Game started for user=%s, session=%s\n", username, gameSessionID)
	}
}
