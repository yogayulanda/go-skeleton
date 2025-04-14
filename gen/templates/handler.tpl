package handler

import (
	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/{{.EntityLower}}"
	"context"
)

type {{.Entity}}Handler struct {
	v1pb.Unimplemented{{.Entity}}ServiceServer
	service *{{.EntityLower}}.{{.Entity}}Service
}

func New{{.Entity}}Handler(service *{{.EntityLower}}.{{.Entity}}Service) *{{.Entity}}Handler {
	return &{{.Entity}}Handler{
		service: service,
	}
}

func (h *{{.Entity}}Handler) Health(ctx context.Context, req *v1pb.HealthRequest) (*v1pb.HealthResponse, error) {
	return h.service.Health(ctx, req)
}
