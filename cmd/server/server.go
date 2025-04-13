package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"
	grpc "gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc"
	grpcgateway "gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc-gateway"
)

func RunServer() {
	// Load configuration
	cfg, err := config.InitConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Initialize DI container
	container := di.InitContainer(cfg)
	logger := container.Logger

	logger.Info("📦 Configuration loaded",
		zap.String("mode", cfg.APP_MODE),
		zap.String("grpc_port", cfg.GRPCPORT),
		zap.String("http_port", cfg.HTTPPORT),
	)

	// Start gRPC server in a separate goroutine
	go func() {
		grpcServer := grpc.NewGRPCServer(container)
		logger.Info("🚀 Starting gRPC server...",
			zap.String("port", cfg.GRPCPORT),
		)
		if err := grpc.StartGRPCServer(cfg.GRPCPORT, grpcServer); err != nil {
			logger.Fatal("❌ Failed to start gRPC server", zap.Error(err))
		}
	}()

	// Start gRPC-Gateway REST proxy in a separate goroutine
	go func() {
		logger.Info("🌐 Starting gRPC-Gateway (REST proxy)...",
			zap.String("port", cfg.HTTPPORT),
		)
		if err := grpcgateway.RunServerGrpcGW(context.Background(), container); err != nil {
			logger.Fatal("❌ Failed to start gRPC-Gateway", zap.Error(err))
		}
	}()

	// Wait for shutdown signal
	waitForShutdown(logger)
}

// waitForShutdown handles OS signals and performs graceful shutdown
func waitForShutdown(logger *zap.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	logger.Info("🛑 Shutdown signal received",
		zap.String("signal", sig.String()),
	)

	logger.Info("✅ Server shutdown completed")
}
