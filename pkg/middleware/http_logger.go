package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ResponseWriter is a custom wrapper for http.ResponseWriter to capture status code
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader method
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// HTTPRequestLogger logs the details of the HTTP request and response
func HTTPRequestLogger(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Wrap the ResponseWriter to capture the status code
			rw := &ResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Get client IP
			clientIP, ClientHost := GetUserIP(r)
			reqID := GetReqID(r)

			// Set request ID & metadata headers for skipping gRPC logging
			r.Header.Set("Grpc-Metadata-X-From-Http", "true")
			r.Header.Set("Grpc-Metadata-X-Request-Id", reqID)

			// Call the next handler
			next.ServeHTTP(rw, r)

			// Log request info
			logger.Info("🌐 HTTP Request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("query", r.URL.RawQuery),
				zap.String("remote_ip", clientIP), // Log real IP
				zap.String("host", ClientHost),
				zap.String("user_agent", r.UserAgent()),
				zap.String("request_id", reqID),
				zap.Int("status_code", rw.statusCode),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}
