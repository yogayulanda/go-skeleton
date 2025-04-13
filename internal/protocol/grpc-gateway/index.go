package grpcgateway

import (
	"context"
	"fmt"
	"net/http"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RunServerGrpcGW(ctx context.Context, container *di.Container) error {
	mux := runtime.NewServeMux()

	// Construct the gRPC server address
	grpcAddress := fmt.Sprintf("localhost:%s", container.Config.GRPCPORT)

	// Register TrxHistoryService handler
	err := v1pb.RegisterTrxHistoryServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddress,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return fmt.Errorf("failed to register TrxHistory handler: %w", err)
	}

	// Register Health handler
	err = v1pb.RegisterHealthHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddress,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return fmt.Errorf("failed to register Health handler: %w", err)
	}

	// Start the HTTP server
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%s", container.Config.HTTPPORT),
		Handler: mux,
	}

	fmt.Printf("gRPC-Gateway server is running on port %s\n", container.Config.HTTPPORT)

	return httpSrv.ListenAndServe()
}
