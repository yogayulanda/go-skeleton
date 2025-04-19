package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// ResponseWriter is a custom wrapper to capture status code.
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader is overridden to capture the status code.
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// HTTPRequestLogger logs incoming HTTP requests and responses.
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

			// Call the next handler

			// If the status code is 4xx or 5xx, log it as an error
			// Log error with 'ERROR' level for 501 Not Implemented
			// Log error with 'WARN' level for 400 Bad Request
			if rw.statusCode == http.StatusBadRequest {
				logger.Warn("400 Bad Request",
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("status_code", fmt.Sprintf("%d", rw.statusCode)),
				)
			}

			// Log error with 'WARN' level for 404 Not Found
			if rw.statusCode == http.StatusNotFound {
				logger.Warn("404 Not Found",
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("status_code", fmt.Sprintf("%d", rw.statusCode)),
				)
			}

			// Log error with 'ERROR' level for 500 Internal Server Error
			if rw.statusCode == http.StatusInternalServerError {
				logger.Error("500 Internal Server Error",
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("status_code", fmt.Sprintf("%d", rw.statusCode)),
				)
			}

			// Log error with 'ERROR' level for 503 Service Unavailable
			if rw.statusCode == http.StatusServiceUnavailable {
				logger.Error("503 Service Unavailable",
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("status_code", fmt.Sprintf("%d", rw.statusCode)),
				)
			}
		})
	}
}

func GetUserIP(r *http.Request) (ip string, host string) {
	host = r.Host

	// Urutan prioritas IP (dari paling dipercaya)
	if ip = r.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip, host
	}
	if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For bisa berisi beberapa IP, ambil IP pertama
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0]), host
	}
	if ip = r.Header.Get("X-Real-IP"); ip != "" {
		return ip, host
	}

	// Fallback ke RemoteAddr
	ip = r.RemoteAddr
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}
	return ip, host
}

func GetReqID(r *http.Request) string {
	reqID := r.Header.Get("X-Request-ID")
	if reqID == "" {
		reqID = uuid.New().String()
	}
	return reqID
}
