package middleware

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryPanicInterceptor returns a unary server interceptor that recovers from panics.
func RecoveryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("🔥 Recovered from panic in gRPC handler",
					zap.Any("panic", r),
					zap.String("method", info.FullMethod),
				)
				err = status.Errorf(codes.Internal, "pkg server error")
			}
		}()
		return handler(ctx, req)
	}
}
