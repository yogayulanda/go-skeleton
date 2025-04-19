package protokol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	v1pb "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGRPCGateway(ctx context.Context, container *di.Container, cfg *config.App) (*http.Server, error) {
	mux := runtime.NewServeMux()

	// Middleware wrapper
	// handler := middleware.ChainMiddleware(
	// 	middleware.HTTPRequestLogger(container.Log),
	// 	middleware.HTTPPanicRecovery(container.Log),
	// 	// Add other middleware here
	// )(mux)

	grpcAddr := fmt.Sprintf("localhost:%s", container.Config.GRPC_PORT)

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := v1pb.RegisterTransactionHistoryServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts)
	if err != nil {
		return nil, err
	}
	err = v1pb.RegisterHealthCheckServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts)
	if err != nil {
		return nil, err
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", container.Config.HTTP_PORT),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return httpServer, nil
}
