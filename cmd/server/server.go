package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/config"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/logging"
	grpc "gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc"
	grpcgateway "gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/protocol/grpc-gateway"
)

func RunServer() {
	// Load configuration
	cfg, err := config.InitConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Initialize logger
	logging.InitLogger(cfg.APP_MODE)
	defer logging.SyncLogger()

	// Initialize DI container
	container := di.InitContainer(cfg)
	log := container.Log

	log.Info("📦 Configuration loaded",
		zap.String("mode", cfg.APP_MODE),
		zap.String("grpc_port", cfg.GRPCPORT),
		zap.String("http_port", cfg.HTTPPORT),
	)

	// Start gRPC server in a separate goroutine
	go func() {
		grpcServer := grpc.NewGRPCServer(container)
		log.Info("🚀 Starting gRPC server...",
			zap.String("port", cfg.GRPCPORT),
		)
		if err := grpc.StartGRPCServer(cfg.GRPCPORT, grpcServer); err != nil {
			log.Fatal("❌ Failed to start gRPC server", zap.Error(err))
		}
	}()

	// Start gRPC-Gateway REST proxy in a separate goroutine
	go func() {
		log.Info("🌐 Starting gRPC-Gateway (REST proxy)...",
			zap.String("port", cfg.HTTPPORT),
		)
		if err := grpcgateway.RunServerGrpcGW(context.Background(), container); err != nil {
			log.Fatal("❌ Failed to start gRPC-Gateway", zap.Error(err))
		}
	}()

	// Wait for shutdown signal
	waitForShutdown(log, container)
}

// waitForShutdown handles OS signals and performs graceful shutdown
func waitForShutdown(log *zap.Logger, container *di.Container) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	log.Info("🛑 Shutdown signal received",
		zap.String("signal", sig.String()),
	)

	// Perform any cleanup or graceful shutdown steps
	// Example: Gracefully shutdown gRPC server or database connections
	// container.DB.Close() // If you need to close DB connections

	// Optionally, wait for ongoing requests to finish before shutting down
	// Timeout to wait for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server gracefully
	// For example, if you're using an HTTP server, you could use:
	// server.Shutdown(ctx)

	log.Info("✅ Server shutdown completed")
}
