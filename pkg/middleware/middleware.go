package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/yogayulanda/go-skeleton/pkg/common"
	"github.com/yogayulanda/go-skeleton/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

// JWTMiddleware untuk validasi token pada HTTP middleware
func HTTPMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ambil token dari header Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Ekstrak token Bearer
			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader { // Jika tidak ada prefix "Bearer "
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Sisipkan token ke dalam konteks request
			ctx := context.WithValue(r.Context(), common.AuthorizationKey, token)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r) // Panggil handler berikutnya
			// Panggil handler berikutnya
		})
	}
}

// rpcMiddleware untuk validasi token pada gRPC middleware
func GrpcMiddleware(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Ambil metadata dari context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("gagal mengambil metadata")
		}
		// Ambil token dari metadata gRPC
		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Error(codes.Unauthenticated, "header Authorization tidak ditemukan")
		}

		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		if tokenString == "" {
			return nil, fmt.Errorf("token tidak ditemukan")
		}

		// Parse dan validasi token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		// Inject claims ke context
		ctx = context.WithValue(ctx, common.CtxUserID, claims.UserID)
		ctx = context.WithValue(ctx, common.CtxRole, claims.Role)

		// Pass ke handler berikutnya
		return handler(ctx, req)
	}
}

// treamGrpcMiddleware untuk validasi token pada gRPC stream middleware
// StreamGrpcMiddleware untuk validasi token pada gRPC stream middleware
func StreamGrpcMiddleware(logger *zap.Logger) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// Ambil metadata dari context
		md, _ := metadata.FromIncomingContext(ss.Context())

		// Ambil token dari metadata gRPC
		token := md.Get("authorization")
		if len(token) == 0 {
			return status.Error(codes.Unauthenticated, "missing Authorization header")
		}

		// Parse dan validasi token
		claims, err := utils.ParseToken(token[0])
		if err != nil {
			return status.Error(codes.Unauthenticated, "invalid token")
		}

		// Inject claims ke context
		ctx := context.WithValue(ss.Context(), common.CtxUserID, claims.UserID)
		ctx = context.WithValue(ctx, common.CtxRole, claims.Role)

		// Gantikan context di ServerStream (dengan context yang sudah diupdate)
		ss = &serverStreamWithContext{
			ServerStream: ss,
			ctx:          ctx,
		}

		// Pass ke stream handler berikutnya
		return handler(srv, ss)
	}
}

// serverStreamWithContext adalah custom ServerStream yang menyimpan context baru
type serverStreamWithContext struct {
	grpc.ServerStream
	ctx context.Context
}

// Context menggantikan context bawaan dengan context baru
func (s *serverStreamWithContext) Context() context.Context {
	return s.ctx
}

// Chain of Unary Interceptors manually
// func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 		// Apply interceptors sequentially
// 		for _, interceptor := range interceptors {
// 			// Wrap the handler to pass to the next interceptor
// 			originalHandler := handler
// 			handler = func(ctx context.Context, req interface{}) (interface{}, error) {
// 				return interceptor(ctx, req, info, originalHandler)
// 			}
// 		}
// 		// Execute the final handler in the chain
// 		return handler(ctx, req)
// 	}
// }
