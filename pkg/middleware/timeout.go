package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TimeoutInterceptor menangani request yang melebihi waktu yang diizinkan (timeout)
func TimeoutInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Mengecek apakah deadline sudah tercapai
		select {
		case <-ctx.Done():
			// Jika deadline tercapai, kembalikan error
			return nil, status.Errorf(codes.DeadlineExceeded, "request timeout")
		default:
			// Lanjutkan dengan handler
			return handler(ctx, req)
		}
	}
}
