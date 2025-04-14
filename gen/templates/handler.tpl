package handler

import (
	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/{{.EntityLower}}"
)

type {{.Entity}}Handler struct {
	v1pb.Unimplemented{{.Entity}}ServiceServer
	Service *{{.EntityLower}}.{{.Entity}}Service
}

func New{{.Entity}}Handler(svc *{{.EntityLower}}.{{.Entity}}Service) *{{.Entity}}Handler {
	return &{{.Entity}}Handler{Service: svc}
}

// TODO: Implement handler methods
