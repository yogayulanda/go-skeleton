package utils

import (
	"github.com/yogayulanda/go-skeleton/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetGrpcError membantu membuat error yang standar dengan code, message, dan metadata
func SetGrpcError(code codes.Code, errCode, message string) error {
	metadata := map[string]string{
		"field":    "id",
		"trace_id": "abc123",
	}
	// Membuat ErrorDetail menggunakan data yang diberikan
	info := &proto.ErrorDetail{
		Code:     errCode,
		Message:  message,
		Metadata: metadata,
	}

	// Membuat status error dengan ErrorDetail
	st := status.New(code, message)
	stWithDetails, _ := st.WithDetails(info)

	return stWithDetails.Err()
}
