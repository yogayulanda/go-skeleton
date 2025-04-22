package protokol

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	v1 "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"
	"github.com/yogayulanda/go-skeleton/pkg/handler"
	"github.com/yogayulanda/go-skeleton/pkg/middleware"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// StartGrpcServer menjalankan gRPC server
func StartGrpcServer(container *di.Container, cfg *config.App) (*grpc.Server, net.Listener, error) {
	// interceptor Panic Custom
	recoveryOpts := []grpc_recovery.Option{
		middleware.PanicRecoveryInterceptor(container.Log),
	}
	// Konfigurasi middleware interceptor (unary dan stream)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),         // Tambahkan tag context
			middleware.LogUnaryInterceptor(container.Log), // Logging dengan zap
			middleware.GrpcMiddleware(container.Log),
			grpc_recovery.UnaryServerInterceptor(),
			middleware.TimeoutInterceptor(),                           // Timeout interceptor
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),     // Recovery dari panic
			apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()), //apmgrpc stracing

		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),             // Untuk stream context tags
			middleware.StreamLoggingInterceptor(container.Log), // Logging stream dengan zap
			middleware.StreamGrpcMiddleware(container.Log),
			grpc_recovery.StreamServerInterceptor(),                    // Recovery stream
			apmgrpc.NewStreamServerInterceptor(apmgrpc.WithRecovery()), //apmgrpc stracing stream
		)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    30 * time.Second,
			Timeout: 10 * time.Second,
		}),
		grpc.MaxConcurrentStreams(1000),
		grpc.MaxRecvMsgSize(4 << 20),
		grpc.MaxSendMsgSize(4 << 20),
	}

	grpcServer := grpc.NewServer(opts...)
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
	reflection.Register(grpcServer)
	return grpcServer, list, nil
}
