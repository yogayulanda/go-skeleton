package protokol

import (
	"fmt"
	"net"

	v1 "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"
	"github.com/yogayulanda/go-skeleton/pkg/handler"
	"google.golang.org/grpc"
)

// StartGrpcServer menjalankan gRPC server
func StartGrpcServer(container *di.Container, cfg *config.App) (*grpc.Server, net.Listener, error) {
	// Setup gRPC server
	// // gRPC server statup options
	// opts := []grpc.ServerOption{
	// 	middleware.Unar/yLoggingInterceptor(container.Log),
	// }

	grpcServer := grpc.NewServer()

	// Setup User handler
	userHandler := handler.NewUserHandler(container.UserService, container.Log)
	v1.RegisterUserServiceServer(grpcServer, userHandler)

	// Setup HealthCheck handler
	healthCheckHandler := handler.NewHealthCheckHandler(container.HealthCheckService, container.Log)
	v1.RegisterHealthCheckServiceServer(grpcServer, healthCheckHandler)

	// Setup listener untuk gRPC server
	list, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC_PORT))
	if err != nil {
		return nil, nil, err // Mengembalikan error jika listener gagal dibuat
	}

	// Register reflection service on gRPC server (untuk debugging)
	// reflection.Register(grpcServer)
	return grpcServer, list, nil
}
