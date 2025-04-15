package {{.EntityLower}}_test

import (
	"testing"

	"github.com/yogayulanda/if-trx-history/internal/domain/{{.EntityLower}}"
)

func TestNew{{.Entity}}Service(t *testing.T) {
	svc := {{.EntityLower}}.New{{.Entity}}Service()
	if svc == nil {
		t.Fatal("expected service to be initialized")
	}
}
