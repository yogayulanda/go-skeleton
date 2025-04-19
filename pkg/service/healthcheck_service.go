package service

import (
	"context"
	"fmt"

	v1 "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HealthCheckService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewHealthCheckService(db *gorm.DB, log *zap.Logger) *HealthCheckService {
	return &HealthCheckService{
		db:  db,
		log: log,
	}
}

func (s *HealthCheckService) CheckHealth(ctx context.Context, req *v1.HealthCheckRequest) (*v1.HealthCheckResponse, error) {
	// Cek koneksi database
	dbStatus := "ok"
	// if err := s.db.DB().PingContext(ctx); err != nil {
	// 	dbStatus = "error"
	// }

	// Cek Kafka connection (misalnya, dengan kafka-go atau librdkafka)
	kafkaStatus := "ok"
	// Kafka connection check logic here...

	// Kembalikan response health check
	return &v1.HealthCheckResponse{
		DbStatus:    dbStatus,
		KafkaStatus: kafkaStatus,
		Status:      "ok", // Status keseluruhan aplikasi
		Message:     fmt.Sprintf("Database: %s, Kafka: %s", dbStatus, kafkaStatus),
	}, nil
}
