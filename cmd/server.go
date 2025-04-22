package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"
	protokol "github.com/yogayulanda/go-skeleton/pkg/protocol"
	"github.com/yogayulanda/go-skeleton/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RunServer starts both the gRPC and HTTP servers, utilizing DI container for dependencies.
func RunServer(container *di.Container, cfg *config.App) error {
	// Mulai gRPC Server, dengan port yang diambil dari config
	grpcServer, list, err := protokol.StartGrpcServer(container, cfg)
	if err != nil {
		container.Log.Error("❌ Failed to start gRPC-server", zap.Error(err))
		return err // Mengembalikan error jika gRPC server gagal dimulai
	}

	// Mulai HTTP server (gRPC Gateway), dengan port yang diambil dari config
	httpServer, err := protokol.StartGRPCGateway(context.Background(), container, cfg)
	if err != nil {
		container.Log.Error("❌ Failed to start gRPC-Gateway", zap.Error(err))
		return err // Mengembalikan error jika HTTP server gagal dimulai
	}

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Mulai gRPC server di goroutine

	go func() {
		container.Log.Info("🌐 Starting gRPC-server", zap.String("port", cfg.GRPC_PORT))
		if err := grpcServer.Serve(list); err != nil {
			container.Log.Fatal("Failed to serve", zap.Any("error", err))
		}
	}()

	go func() {
		container.Log.Info("🌐 Starting HTTP server for gRPC Gateway", zap.String("port", cfg.HTTP_PORT))
		if err := httpServer.ListenAndServe(); err != nil {
			// handleHTTPServerShutdownError(err, container)
		}
	}()

	// Log service started AFTER all init
	utils.LogAvailableEndpoints()
	container.Log.Info("✅ go-skeleton service started successfully",
		zap.String("version", "v1.0.0"),
		zap.String("time", time.Now().Format(time.RFC3339)),
	)
	// Menunggu signal untuk shutdown
	<-stop
	container.Log.Info("🛑 Received shutdown signal")
	// Menjalankan graceful shutdown untuk gRPC server
	gracefulShutdown(grpcServer, httpServer, container)
	return nil
}

// gracefulShutdown menangani shutdown gRPC dan HTTP Gateway secara bersamaan
func gracefulShutdown(grpcServer *grpc.Server, httpServer *http.Server, container *di.Container) {
	// Menjalankan graceful shutdown untuk gRPC server
	grpcServer.GracefulStop()
	container.Log.Info("🌐 gRPC server stopped gracefully")

	// Menjalankan graceful shutdown untuk HTTP Gateway
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		handleHTTPServerShutdownError(err, container)
	}
	container.Log.Info("✅ Server shutdown completed")
}

// handleHTTPServerShutdownError menangani error shutdown server HTTP dengan lebih terorganisir
func handleHTTPServerShutdownError(err error, container *di.Container) {
	if err == http.ErrServerClosed {
		// Server sudah ditutup secara graceful
		container.Log.Info("🌐 HTTP Gateway closed gracefully")
	} else {
		// Server gagal dijalankan
		container.Log.Fatal("Failed to serve HTTP gateway", zap.Error(err))
	}
}
