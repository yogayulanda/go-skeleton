package grpcgateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/di"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/middleware"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/utils"
	"go.uber.org/zap"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunServerGrpcGW(ctx context.Context, container *di.Container) error {
	mux := runtime.NewServeMux()

	// Middleware wrapper
	handler := middleware.ChainMiddleware(
		middleware.HTTPRequestLogger(container.Logger),
		middleware.HTTPPanicRecovery(container.Logger),
		// Add other middleware here
	)(mux)

	grpcAddr := fmt.Sprintf("localhost:%s", container.Config.GRPCPORT)

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := v1pb.RegisterTrxHistoryServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
		return fmt.Errorf("failed to register TrxHistory handler: %w", err)
	}
	if err := v1pb.RegisterHealthHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
		return fmt.Errorf("failed to register Health handler: %w", err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", container.Config.HTTPPORT),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Optional TLS setup if enabled
	if container.Config.ENABLE_TLS {
		srv.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	utils.LogAvailableEndpoints(container.Logger)
	// Log service started AFTER all init
	container.Logger.Info("✅ if-trx-history service started successfully",
		zap.String("version", "v1.0.0"),
		zap.String("time", time.Now().Format(time.RFC3339)),
	)

	// Graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop

		container.Logger.Info("🛑 Shutting down HTTP server...")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctxTimeout); err != nil {
			container.Logger.Error("❌ Failed to shutdown HTTP server gracefully", zap.Error(err))
		}
		close(idleConnsClosed)
	}()

	// Run server
	var err error
	if container.Config.ENABLE_TLS {
		err = srv.ListenAndServeTLS(container.Config.TLS_CERT_PATH, container.Config.TLS_KEY_PATH)
	} else {
		err = srv.ListenAndServe()
	}

	if err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	<-idleConnsClosed
	return nil
}
