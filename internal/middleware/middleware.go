package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
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

// Chain of Unary Interceptors manually
func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Apply interceptors sequentially
		for _, interceptor := range interceptors {
			// Wrap the handler to pass to the next interceptor
			originalHandler := handler
			handler = func(ctx context.Context, req interface{}) (interface{}, error) {
				return interceptor(ctx, req, info, originalHandler)
			}
		}
		// Execute the final handler in the chain
		return handler(ctx, req)
	}
}
