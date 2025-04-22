package common

// CtxKey adalah tipe kunci untuk context, menggunakan custom type dari string
type CtxKey string

// Deklarasi constant untuk key-context
const (
	CtxUserID        CtxKey = "user_id"
	CtxRole          CtxKey = "role"
	CtxTraceID       CtxKey = "trace_id"
	AuthorizationKey CtxKey = "Authorization"
)
