package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"

	"github.com/yogayulanda/go-skeleton/pkg/common"
	"go.uber.org/zap"
)

// ErrorResponse untuk format JSON HTTP error response
type ErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// CustomHTTPErrorHandler menangani error gRPC dan konversi ke JSON standar HTTP
func CustomHTTPErrorHandler(
	ctx context.Context,
	mux *runtime.ServeMux,
	m runtime.Marshaler,
	w http.ResponseWriter,
	r *http.Request,
	err error,
	log *zap.Logger,
) {
	st, ok := status.FromError(err)
	if !ok {
		http.Error(w, "Unknown error", http.StatusInternalServerError)
		return
	}

	// Ambil errorKey dari st.Message()
	errorKey := st.Message()

	// Mapping error code dan message dari ErrorMap atau Redis
	mapped := common.GetErrorEntry(errorKey) // Hanya 1 kali dipanggil

	// Tentukan HTTP status code berdasarkan gRPC error code
	statusCode := runtime.HTTPStatusFromCode(st.Code())

	// Kirimkan response error dalam format JSON
	handleError(w, mapped.Code, mapped.Message, statusCode, log)
}

func handleError(w http.ResponseWriter, code string, message string, statusCode int, log *zap.Logger) {
	// Kirimkan response error dalam format JSON
	resp := ErrorResponse{
		Status:  "error",
		Code:    code,
		Message: message,
	}
	writeJSON(w, statusCode, resp)
}

func writeJSON(w http.ResponseWriter, statusCode int, resp ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}
