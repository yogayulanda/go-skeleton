package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.elastic.co/apm" // pastikan ini benar, bukan "elastic.co/apm"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func LogUnaryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()

		// Ambil trace_id dari APM jika ada
		traceID := getTraceIDFromAPM(ctx)
		requestID := ""

		// Ambil metadata dari request
		md, _ := metadata.FromIncomingContext(ctx)

		// Ambil request_id dari metadata (atau generate jika kosong)
		if vals := md.Get("x-request-id"); len(vals) > 0 {
			requestID = vals[0]
		} else {
			requestID = uuid.New().String()
		}

		// Jika request datang dari HTTP Gateway, skip log request di gRPC
		if len(md.Get("x-from-http")) > 0 {
			// Dapatkan trace_id dari HTTP Gateway
			traceID = md.Get("apm-trace-id")[0]
			// Skip log request karena sudah ada log di HTTP
			logger.Debug("Skipping gRPC request log as it's already logged in HTTP Gateway")
		}

		// Log request hanya jika request belum diproses di HTTP Gateway
		if len(md.Get("x-from-http")) == 0 {
			logger.Info("📡 gRPC Unary Request",
				zap.String("method", info.FullMethod),
				zap.String("trace_id", traceID),
				zap.String("request_id", requestID),
				zap.String("request", fmt.Sprintf("%v", req)),
			)
		}

		// Eksekusi gRPC handler
		resp, err := handler(ctx, req)
		code := status.Code(err)

		// Log error jika ada
		if err != nil {
			if code == codes.Unimplemented {
				logger.Warn("⚠️ Method Unimplemented",
					zap.String("method", info.FullMethod),
					zap.String("trace_id", traceID),
					zap.String("request_id", requestID),
					zap.String("status_code", code.String()),
					zap.String("error", err.Error()),
				)
			} else {
				logger.Error("❌ RPC Error",
					zap.String("method", info.FullMethod),
					zap.String("trace_id", traceID),
					zap.String("request_id", requestID),
					zap.String("status_code", code.String()),
					zap.Error(err),
				)
			}
		}

		// Log response selalu dicatat
		logger.Info("✅ gRPC Unary Response",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.String("request_id", requestID),
			zap.String("status_code", code.String()),
			zap.Duration("duration", time.Since(start)),
		)

		return resp, err
	}
}

func StreamLoggingInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()
		ctx := ss.Context()

		// Ambil trace_id dari APM
		traceID := getTraceIDFromAPM(ctx)

		// Ambil request_id dari metadata
		requestID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if vals := md.Get("x-request-id"); len(vals) > 0 {
				requestID = vals[0]
			}
		}

		logger.Info("📡 gRPC Stream Request",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.String("request_id", requestID),
			zap.Bool("is_client_stream", info.IsClientStream),
			zap.Bool("is_server_stream", info.IsServerStream),
		)

		err := handler(srv, ss)
		code := status.Code(err)
		tags := grpc_ctxtags.Extract(ctx).Values()

		fields := []zap.Field{
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.String("request_id", requestID),
			zap.Duration("duration", time.Since(start)),
			zap.Any("tags", tags),
			zap.String("grpc_code", code.String()),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
			logger.Error("❌ gRPC Stream Failed", fields...)
		} else {
			logger.Info("✅ gRPC Stream Completed", fields...)
		}

		return err
	}
}

// getTraceIDFromAPM membantu mendapatkan trace_id dari context APM jika ada
func getTraceIDFromAPM(ctx context.Context) string {
	// Ambil span dari context
	span := apm.SpanFromContext(ctx)
	if span != nil {
		// Jika ada span, ambil trace ID-nya
		return span.TraceContext().Trace.String()
	}
	return ""
}
