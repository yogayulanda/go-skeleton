package protokol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	v1pb "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/config"
	"github.com/yogayulanda/go-skeleton/pkg/di"
	"github.com/yogayulanda/go-skeleton/pkg/middleware"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp" // ✅ Tambahkan ini
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func StartGRPCGateway(ctx context.Context, container *di.Container, cfg *config.App) (*http.Server, error) {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(middleware.CustomHTTPErrorHandler), // 🛠️ Custom Error Mapping

		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			md := metadata.New(nil)

			// Inject request ID
			reqID := r.Header.Get("X-Request-ID")
			if reqID != "" {
				md.Set("x-request-id", reqID)
			}

			// Inject APM trace info
			if tx := apm.TransactionFromContext(r.Context()); tx != nil {
				tc := tx.TraceContext()
				md.Set("apm-trace-id", tc.Trace.String())
				md.Set("apm-span-id", tc.Span.String())
			}

			md.Set("x-from-http", "true")
			return md
		}),
	)

	// ✅ Bungkus middleware seperti biasa
	handler := middleware.ChainMiddleware(
		middleware.HTTPRequestLogger(container.Log),
		middleware.HTTPPanicRecovery(container.Log),
		middleware.HTTPMiddleware(container.Log), // Add JWT middleware
	)(mux)

	// ✅ Bungkus dengan APM agar transaction aktif di context
	wrappedHandler := apmhttp.Wrap(handler)

	grpcAddr := fmt.Sprintf("localhost:%s", container.Config.GRPC_PORT)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := v1pb.RegisterHealthCheckServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
		return nil, fmt.Errorf("failed to register HealthCheckService handler: %v", err)
	}

	if err := v1pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
		return nil, fmt.Errorf("failed to register UserService handler: %v", err)
	}

	// Gunakan wrappedHandler di sini
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", container.Config.HTTP_PORT),
		Handler:      wrappedHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return httpServer, nil
}
