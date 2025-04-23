package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/yogayulanda/go-skeleton/pkg/common"
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
) {
	st, ok := status.FromError(err)
	if !ok {
		http.Error(w, "Unknown error", http.StatusInternalServerError)
		return
	}

	// Tangani endpoint tidak ditemukan
	if st.Code() == codes.Unimplemented || st.Code() == codes.NotFound {
		resp := ErrorResponse{
			Status:  "error",
			Code:    "E404",
			Message: "Endpoint tidak ditemukan",
		}
		writeJSON(w, http.StatusNotFound, resp)
		return
	}

	// Tangani unauthorized
	if st.Code() == codes.Unauthenticated {
		resp := ErrorResponse{
			Status:  "error",
			Code:    "E401",
			Message: "Anda belum login atau token tidak valid",
		}
		writeJSON(w, http.StatusUnauthorized, resp)
		return
	}

	// Tangani forbidden
	if st.Code() == codes.PermissionDenied {
		resp := ErrorResponse{
			Status:  "error",
			Code:    "E403",
			Message: "Anda tidak memiliki akses ke resource ini",
		}
		writeJSON(w, http.StatusForbidden, resp)
		return
	}

	// Tangani input error
	if st.Code() == codes.InvalidArgument {
		resp := ErrorResponse{
			Status:  "error",
			Code:    "E400",
			Message: "Input tidak valid",
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	// Mapping dari ErrorMap (error business logic)
	errorKey := st.Message()
	mapped := common.GetErrorEntry(errorKey)

	resp := ErrorResponse{
		Status:  "error",
		Code:    mapped.Code,
		Message: mapped.Message,
	}
	writeJSON(w, runtime.HTTPStatusFromCode(st.Code()), resp)
}

func writeJSON(w http.ResponseWriter, statusCode int, resp ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(resp)
}
