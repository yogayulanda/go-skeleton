package handler

import (
	"context"
	"internal/domain/{{.ServicePackage}}"
	pb "{{.ProtoImportPath}}"
)

type {{.ServiceName}}Handler struct {
	service {{.ServicePackage}}.Service
}

func New{{.ServiceName}}Handler(s {{.ServicePackage}}.Service) *{{.ServiceName}}Handler {
	return &{{.ServiceName}}Handler{service: s}
}

{{range .Methods}}
func (h *{{$.ServiceName}}Handler) {{.Name}}(ctx context.Context, req *pb.{{.InputType}}) (*pb.{{.OutputType}}, error) {
	return h.service.{{.Name}}(ctx, req)
}
{{end}}
