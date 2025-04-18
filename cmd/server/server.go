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
func RunServer(appContext *di.Container, cfg *config.App) {
	// Start gRPC server in a separate goroutine
	go startGRPCServer(appContext, cfg)

	// Start gRPC-Gateway REST proxy in a separate goroutine
	go startGRPCGateway(appContext, cfg)

	// Wait for shutdown signal
	waitForShutdown(appContext.Log)
}

// startGRPCServer initializes and runs the gRPC server
func startGRPCServer(container *di.Container, cfg *config.App) {
	grpcServer := grpc.NewGRPCServer(container)
	container.Log.Info("🚀 Starting gRPC server...", zap.String("port", cfg.GRPC_PORT))

	if err := grpc.StartGRPCServer(cfg.GRPC_PORT, grpcServer); err != nil {
		container.Log.Fatal("❌ Failed to start gRPC server", zap.Error(err))
	}
}

// startGRPCGateway initializes and runs the gRPC-Gateway (REST proxy)
func startGRPCGateway(container *di.Container, cfg *config.App) {
	container.Log.Info("🌐 Starting gRPC-Gateway (REST proxy)...", zap.String("port", cfg.HTTP_PORT))

	if err := grpcgateway.RunServerGrpcGW(context.Background(), container); err != nil {
		container.Log.Fatal("❌ Failed to start gRPC-Gateway", zap.Error(err))
	}
}

// waitForShutdown handles graceful shutdown of the application
func waitForShutdown(log *zap.Logger) {
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
