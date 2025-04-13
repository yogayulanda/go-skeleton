package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LoggingMiddleware logs details of every HTTP request, including the endpoint
func LoggingMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Log request details
			logger.Info("Incoming HTTP request",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("path", r.URL.Path),
				zap.String("remote_addr", r.RemoteAddr),
			)

			// Call the next handler
			next.ServeHTTP(w, r)

			// Log response duration
			duration := time.Since(start)
			logger.Info("HTTP request completed",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", duration),
			)
		})
	}
}
