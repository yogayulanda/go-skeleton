package middleware

import (
	"context"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryPanicInterceptor returns a unary server interceptor that recovers from panics.
// PanicRecoveryInterceptor returns recovery options with zap logging
func PanicRecoveryInterceptor(logger *zap.Logger) grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandlerContext(func(ctx context.Context, p interface{}) error {
		logger.Error("🔥 Panic Recovered",
			zap.Any("panic", p),
		)
		return status.Errorf(codes.Internal, "internal server error")
	})
}
