package {{.EntityLower}}

type {{.Entity}}Service struct{}

func New{{.Entity}}Service() *{{.Entity}}Service {
	return &{{.Entity}}Service{}
}

// TODO: Implement business logic
