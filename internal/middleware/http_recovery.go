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
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
