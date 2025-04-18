package grpc

import (
	"fmt"
	"net"

	v1pb "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/internal/di"
	"github.com/yogayulanda/go-skeleton/internal/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	v1pb.RegisterUserServiceServer(grpcServer, container.UserHandler)
	v1pb.RegisterTransactionHistoryServiceServer(grpcServer, container.TrxHandler)
	v1pb.RegisterHealthServiceServer(grpcServer, container.HealthHandler)
	// v1pb.RegisterHealthServiceServer(grpcServer, container.HealthHandler)
	reflection.Register(grpcServer)
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
