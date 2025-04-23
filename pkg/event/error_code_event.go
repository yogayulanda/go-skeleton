package event

import (
	"context"
	"time"

	redisClient "github.com/redis/go-redis/v9"
	"github.com/yogayulanda/go-skeleton/pkg/redis"
	"github.com/yogayulanda/go-skeleton/pkg/repository"
	"go.uber.org/zap"
)

type ErrorCodeEvent struct {
	Cache  *redis.ErrorCache
	Logger *zap.Logger
	Repo   *repository.ErrorCodeRepository
}

func NewErrorCodeEvent(cache *redis.ErrorCache, repo *repository.ErrorCodeRepository, logger *zap.Logger) *ErrorCodeEvent {
	return &ErrorCodeEvent{
		Cache:  cache,
		Logger: logger,
		Repo:   repo,
	}
}

// Subscribe mendengarkan perubahan error code dan memperbarui cache Redis
func (e *ErrorCodeEvent) SubscribeErrorCodeChanges(ctx context.Context, client *redisClient.Client, channel string) {
	pubsub := client.Subscribe(ctx, channel)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		if msg.Payload == "update" {
			// Update cache hanya untuk error code yang diubah
			e.Logger.Info("Error code update detected. Refreshing cache.")
			e.UpdateErrorCodeCache(ctx)
		}
	}
}

// UpdateErrorCodeCache mengambil error code yang diubah dan memperbarui Redis
func (e *ErrorCodeEvent) UpdateErrorCodeCache(ctx context.Context) {
	// Ambil timestamp saat ini untuk mendeteksi perubahan error code
	since := time.Now().Add(-24 * time.Hour) // Ambil error codes yang diubah dalam 24 jam terakhir

	// Ambil semua error codes yang diubah
	errorCodes, err := e.Repo.GetUpdatedErrorCodes(ctx, since)
	if err != nil {
		e.Logger.Error("Failed to get updated error codes from DB", zap.Error(err))
		return
	}

	// Update Redis untuk setiap error code yang diubah
	for _, errorCode := range errorCodes {
		ttl := 24 * time.Hour
		err = e.Cache.SetErrorCode(ctx, errorCode.Code, errorCode.Message, ttl)
		if err != nil {
			e.Logger.Error("Failed to update error code cache", zap.String("code", errorCode.Code), zap.Error(err))
		}
	}
}
