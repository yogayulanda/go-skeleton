package handler

import (
	"context"

	v1 "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/service"
	"go.uber.org/zap"
)

type HealthCheckHandler struct {
	v1.UnimplementedHealthCheckServiceServer // Embedkan UnimplementedUserServiceServer untuk kompatibilitas gRPC
	healthService                            *service.HealthCheckService
	log                                      *zap.Logger
}

func NewHealthCheckHandler(healthService *service.HealthCheckService, log *zap.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{
		healthService: healthService,
		log:           log,
	}
}

// CheckHealth menangani permintaan untuk memeriksa status kesehatan aplikasi
func (h *HealthCheckHandler) CheckHealth(ctx context.Context, req *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	h.log.Info("Handling HealthCheck request")

	// Memanggil HealthCheckService untuk memeriksa status
	healthStatus, err := h.healthService.CheckHealth(ctx, req)
	if err != nil {
		h.log.Error("Error checking health", zap.Error(err))
		return nil, err
	}

	// Mengembalikan response dengan status kesehatan aplikasi
	return &v1.HealthCheckResponse{
		DbStatus:    healthStatus.DbStatus,
		KafkaStatus: healthStatus.KafkaStatus,
		Status:      healthStatus.Status,
		Message:     healthStatus.Message,
	}, nil
}
