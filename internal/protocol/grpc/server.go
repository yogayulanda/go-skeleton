package grpc

import (
	"fmt"
	"net"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/middleware"

	"google.golang.org/grpc"
)

// NewGRPCServer creates a new gRPC server with chained interceptors
func NewGRPCServer(container *di.Container) *grpc.Server {
	// Manually chain unary interceptors
	chainUnary := middleware.ChainUnaryServer(
		middleware.UnaryPanicInterceptor(container.Log),   // Panic recovery interceptor
		middleware.UnaryLoggingInterceptor(container.Log), // Logging interceptor
	)

	// Stream interceptor (for stream requests)
	// chainStream := StreamLoggingInterceptor(container.Logger)

	// Create the gRPC server with interceptors
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(chainUnary),
		// grpc.StreamInterceptor(chainStream),
	)

	// Register the gRPC service handlers
	v1pb.RegisterTrxHistoryServiceServer(grpcServer, container.TrxHandler)
	v1pb.RegisterHealthServer(grpcServer, container.HealthHandler)

	// @auto:inject:handler

	return grpcServer
}

func StartGRPCServer(port string, server *grpc.Server) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}
	return server.Serve(listener)
}
