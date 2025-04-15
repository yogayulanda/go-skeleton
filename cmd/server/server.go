package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yogayulanda/go-skeleton/internal/config"
	"github.com/yogayulanda/go-skeleton/internal/di"
	grpc "github.com/yogayulanda/go-skeleton/internal/protocol/grpc"
	grpcgateway "github.com/yogayulanda/go-skeleton/internal/protocol/grpc-gateway"
	"go.uber.org/zap"
)

// RunServer starts both the gRPC and HTTP servers, utilizing DI container for dependencies.
func RunServer(container *di.Container, log *zap.Logger, cfg *config.App) {
	// Start gRPC server in a separate goroutine
	go startGRPCServer(container, log, cfg)

	// Start gRPC-Gateway REST proxy in a separate goroutine
	go startGRPCGateway(container, log, cfg)

	// Wait for shutdown signal
	waitForShutdown(log, container)
}

// startGRPCServer initializes and runs the gRPC server
func startGRPCServer(container *di.Container, log *zap.Logger, cfg *config.App) {
	grpcServer := grpc.NewGRPCServer(container)
	log.Info("🚀 Starting gRPC server...", zap.String("port", cfg.GRPC_PORT))

	if err := grpc.StartGRPCServer(cfg.GRPC_PORT, grpcServer); err != nil {
		log.Fatal("❌ Failed to start gRPC server", zap.Error(err))
	}
}

// startGRPCGateway initializes and runs the gRPC-Gateway (REST proxy)
func startGRPCGateway(container *di.Container, log *zap.Logger, cfg *config.App) {
	log.Info("🌐 Starting gRPC-Gateway (REST proxy)...", zap.String("port", cfg.HTTP_PORT))

	if err := grpcgateway.RunServerGrpcGW(context.Background(), container); err != nil {
		log.Fatal("❌ Failed to start gRPC-Gateway", zap.Error(err))
	}
}

// waitForShutdown handles graceful shutdown of the application
func waitForShutdown(log *zap.Logger, container *di.Container) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	log.Info("🛑 Shutdown signal received", zap.String("signal", sig.String()))

	// Perform cleanup or graceful shutdown steps
	// Example: Gracefully shutdown gRPC server or close database connections
	// container.DB.Close() // Uncomment if needed

	// Optionally, wait for ongoing requests to finish before shutting down
	// Timeout to wait for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("✅ Server shutdown completed")
}
