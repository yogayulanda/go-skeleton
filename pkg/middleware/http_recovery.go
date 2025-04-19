package middleware

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
)

// HTTPPanicRecovery returns a middleware that recovers from panics.
func HTTPPanicRecovery(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rcv := recover(); rcv != nil {
					logger.Error("🔥 Recovered from panic in HTTP handler",
						zap.Any("panic", rcv),
						zap.String("url", r.URL.String()),
						zap.String("method", r.Method),
						zap.ByteString("stack", debug.Stack()),
					)
					http.Error(w, "pkg Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func AddRecoveryMiddleware(next http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer to recover from panic
		defer func() {
			if r := recover(); r != nil {
				// Memastikan bahwa `r` adalah tipe yang tepat (*http.Request)
				if req, ok := r.(*http.Request); ok {
					// Log panic with ERROR level
					logger.Error("Panic recovered",
						zap.String("method", req.Method),
						zap.String("url", req.URL.String()),
						zap.String("remote_addr", req.RemoteAddr),
						zap.Any("panic", r),
					)
					// Send internal server error response
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}
		}()

		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}
