package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// InitRedis — создаёт клиента Redis
func InitRedis(addr, password string, dbNumber int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbNumber,
	})

	// Быстрый тест (PING)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}
	log.Printf("Connected to Redis (%s)\n", addr)
	return rdb, nil
}

// TestRedisConnection — пример: запись-чтение тестового ключа
func TestRedisConnection(ctx context.Context, rdb *redis.Client) error {
	key := "testKey"
	value := "Hello from Redis!"
	err := rdb.Set(ctx, key, value, 10*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("redis SET error: %w", err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("redis GET error: %w", err)
	}
	log.Println("Redis testKey =", val)
	return nil
}
