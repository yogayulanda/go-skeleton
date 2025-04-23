package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

func InitRedist(cfg *App, log *zap.Logger) (*redis.Client, error) {
	if redisClient != nil {
		// Redis client sudah diinisialisasi
		log.Info("Redis client already initialized.")
		return redisClient, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIST_ADDR,     // e.g. "localhost:6379"
		Password: cfg.REDIST_PASSWORD, // leave empty if not using password
		DB:       cfg.REDIST_DB,       // e.g. 0
	})

	// Cek koneksi Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	log.Info("Redis successfully init")
	return client, nil
}

// GetRedisClient mengembalikan Redis client yang sudah diinisialisasi
func GetRedisClient() *redis.Client {
	return redisClient
}
