package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TimeoutInterceptor sets timeout per RPC method
func TimeoutInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var timeout time.Duration

		// Sesuaikan timeout per method
		switch info.FullMethod {
		case "/proto.v1.FinanceService/CheckBill": //set timeout untuk endpoint ke partner
			timeout = 10 * time.Second
		case "/proto.v1.SyncCatalogService/Sync":
			timeout = 15 * time.Second
		default:
			timeout = 5 * time.Second // default global timeout
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		ch := make(chan struct{})
		var resp interface{}
		var err error

		go func() {
			resp, err = handler(ctx, req)
			close(ch)
		}()

		select {
		case <-ctx.Done():
			return nil, status.Error(codes.DeadlineExceeded, "request timeout")
		case <-ch:
			return resp, err
		}
	}
}
