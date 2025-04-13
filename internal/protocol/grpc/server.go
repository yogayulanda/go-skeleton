package grpc

import (
	"fmt"
	"net"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"

	"google.golang.org/grpc"
)

func NewGRPCServer(container *di.Container) *grpc.Server {
	server := grpc.NewServer()

	// Register the handlers to the server
	v1pb.RegisterTrxHistoryServiceServer(server, container.TrxHandler)
	v1pb.RegisterHealthServer(server, container.HealthHandler)

	return server
}

func StartGRPCServer(port string, server *grpc.Server) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}

	fmt.Printf("gRPC server is running on port %s\n", port)
	return server.Serve(listener)
}
