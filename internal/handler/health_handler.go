package handler

import (
	"context"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
)

// TrxHistoryHandler mengimplementasikan v1pb.TrxHistoryServiceServer
type HealthHandler struct {
	v1pb.UnimplementedHealthServer
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(ctx context.Context, _ *v1pb.HealthCheckRequest) (*v1pb.HealthCheckResponse, error) {
	// Simulasi pengecekan status dependencies
	statuses := map[string]string{
		"db":    "OK",
		"kafka": "OK",
		"redis": "OK",
	}
	return &v1pb.HealthCheckResponse{
		HealthStatus:      "SERVING",
		ComponentStatuses: statuses,
	}, nil
}
