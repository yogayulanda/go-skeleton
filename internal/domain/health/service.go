package health

import (
	"context"

	v1proto "github.com/yogayulanda/go-skeleton/gen/proto/v1" // Import for v1 Protobuf
	// Import for v2 Protobuf (if applicable)

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Service struct {
	redisClient *redis.Client
	kafkaReader *kafka.Reader
	logger      *zap.Logger
}

// NewService creates a new instance of the Health service
func NewService(redisClient *redis.Client, kafkaReader *kafka.Reader, logger *zap.Logger) *Service {
	return &Service{
		redisClient: redisClient,
		kafkaReader: kafkaReader,
		logger:      logger,
	}
}

// Check implements the HealthServer interface for v1
func (s *Service) CheckHealth(ctx context.Context, req *v1proto.HealthCheckRequest) (*v1proto.HealthCheckResponse, error) {
	componentStatuses := make(map[string]string)

	// Check Redis health
	if _, err := s.redisClient.Ping(ctx).Result(); err != nil {
		componentStatuses["redis"] = "FAIL"
		s.logger.Warn("Redis health check failed", zap.Error(err))
	} else {
		componentStatuses["redis"] = "OK"
	}

	// Check Kafka health
	if err := s.kafkaReader.SetOffset(kafka.LastOffset); err != nil {
		componentStatuses["kafka"] = "FAIL"
		s.logger.Warn("Kafka health check failed", zap.Error(err))
	} else {
		componentStatuses["kafka"] = "OK"
	}

	// Simulate database health check (always OK for now)
	componentStatuses["database"] = "OK"

	// Determine overall health status
	overallStatus := "SERVING"
	for _, status := range componentStatuses {
		if status == "FAIL" {
			overallStatus = "NOT_SERVING"
			break
		}
	}

	return &v1proto.HealthCheckResponse{
		HealthStatus:      overallStatus,
		ComponentStatuses: componentStatuses,
	}, nil
}

// Add v2 implementation here if applicable
// Uncomment and implement the following if v2 is needed
// var _ v2proto.HealthServer = (*Service)(nil)
// func (s *Service) CheckV2(ctx context.Context, req *v2proto.HealthCheckRequest) (*v2proto.HealthCheckResponse, error) {
//     // Implement v2 logic here
// }
