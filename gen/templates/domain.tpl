package {{.EntityLower}}

import (
	"context"
	v1pb "github.com/yogayulanda/go-skeleton/api/proto/gen/v1"
)

type {{.Entity}}Service struct {
	// Add dependencies (e.g. repository, logger) here if needed
}

func New{{.Entity}}Service() *{{.Entity}}Service {
	return &{{.Entity}}Service{}
}

func (s *{{.Entity}}Service) Health(ctx context.Context, req *v1pb.HealthRequest) (*v1pb.HealthResponse, error) {
    // Simulate health check logic
	componentStatuses := make(map[string]string)
	componentStatuses["redis"] = "OK"

	return &v1pb.HealthResponse{
	    HealthStatus:      "OK",
		ComponentStatuses: componentStatuses, // You can generate real timestamp here
	}, nil
}
