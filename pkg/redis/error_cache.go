package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogayulanda/go-skeleton/pkg/repository"
	"go.uber.org/zap"
)

// ErrorCache untuk menyimpan dan mengelola error code di Redis
type ErrorCache struct {
	Client *redis.Client
	Logger *zap.Logger
}

// InitErrorCache menginisialisasi ErrorCache dengan Redis client
func InitErrorCache(client *redis.Client, logger *zap.Logger) *ErrorCache {
	return &ErrorCache{
		Client: client,
		Logger: logger,
	}
}

// SetErrorCode menyimpan error code ke Redis dengan TTL
func (e *ErrorCache) SetErrorCode(ctx context.Context, code string, message string, ttl time.Duration) error {
	key := fmt.Sprintf("error_code:%s", code)
	err := e.Client.Set(ctx, key, message, ttl).Err()
	if err != nil {
		e.Logger.Error("failed to set error code cache", zap.String("code", code), zap.Error(err))
		return err
	}
	return nil
}

// GetErrorCode mengambil error code dari Redis, jika tidak ada baru ambil dari DB
func (e *ErrorCache) GetErrorCode(ctx context.Context, code string, repo *repository.ErrorCodeRepository) (string, error) {
	key := fmt.Sprintf("error_code:%s", code)
	val, err := e.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Cache miss, tidak ada di Redis, ambil dari DB
		message, err := repo.GetErrorCode(ctx, code)
		if err != nil {
			e.Logger.Error("failed to get error code from DB", zap.String("code", code), zap.Error(err))
			return "", err
		}
		// Simpan ke Redis
		ttl := 24 * time.Hour // Set TTL ke 24 jam
		err = e.SetErrorCode(ctx, code, message, ttl)
		if err != nil {
			return "", err
		}
		return message, nil
	}
	if err != nil {
		e.Logger.Error("failed to get error code cache", zap.String("code", code), zap.Error(err))
		return "", err
	}
	return val, nil
}
