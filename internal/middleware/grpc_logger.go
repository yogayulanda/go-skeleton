package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryLoggingInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		start := time.Now()
		// Cek x-from-http untuk skip logging request pada HTTP
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if val := md.Get("x-from-http"); len(val) > 0 && val[0] == "true" {
				//get traceID from x-request-id
				traceID := ""
				if vals := md.Get("x-request-id"); len(vals) > 0 {
					traceID = vals[0]
				} else {
					traceID = uuid.New().String()
				}
				// Log hanya response-nya, tanpa request
				resp, err := handler(ctx, req)
				// Log the response details
				code := status.Code(err)
				logger.Info("📡 gRPC Unary Response for HTTP",
					zap.String("method", info.FullMethod),
					zap.String("trace_id", traceID),
					zap.Duration("duration", time.Since(start)), // Record duration of request
					zap.String("status_code", code.String()),
					zap.Error(err),
				)
				return resp, err
			}
		}

		// Jika tidak ada metadata x-from-http, log request dan response seperti biasa

		traceID := getOrGenerateTraceID(ctx)

		// Log incoming request
		logger.Info("📡 gRPC Unary Request",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.String("request", fmt.Sprintf("%v", req)),
		)

		// Execute the gRPC handler
		resp, err := handler(ctx, req)

		// Log the response details
		code := status.Code(err)
		logger.Info("📡 gRPC Unary Response",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.Duration("duration", time.Since(start)),
			zap.String("status_code", code.String()),
			zap.Error(err),
		)

		return resp, err
	}
}

func getOrGenerateTraceID(ctx context.Context) string {
	md, _ := metadata.FromIncomingContext(ctx)
	if traceIDs, ok := md["x-trace-id"]; ok && len(traceIDs) > 0 {
		return traceIDs[0] // Use trace ID from incoming metadata if exists
	}
	// If no trace ID, generate a new one
	return uuid.New().String()
}

func StreamLoggingInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		start := time.Now()

		// Generate or retrieve Trace ID from metadata
		traceID := getOrGenerateTraceID(ss.Context())

		// Log the incoming request
		logger.Info("📡 gRPC Stream Request",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.Bool("is_client_stream", info.IsClientStream),
			zap.Bool("is_server_stream", info.IsServerStream),
		)

		err := handler(srv, ss)
		code := status.Code(err)

		// Log the response details
		logger.Info("📡 gRPC Stream Response",
			zap.String("method", info.FullMethod),
			zap.String("trace_id", traceID),
			zap.Duration("duration", time.Since(start)),
			zap.String("status_code", code.String()),
			zap.Error(err),
		)

		return err
	}
}
