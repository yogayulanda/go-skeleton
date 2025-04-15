package handler

import (
	"context"

	v1pb "github.com/yogayulanda/if-trx-history/gen/proto/v1"
	"github.com/yogayulanda/if-trx-history/internal/domain/history"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TrxHistoryHandler struct {
	v1pb.UnimplementedTransactionHistoryServiceServer // embed gRPC compatibility
	service                                           *history.TrxHistoryService
}

func NewTrxHistoryHandler(service *history.TrxHistoryService) *TrxHistoryHandler {
	return &TrxHistoryHandler{service: service}
}

func (h *TrxHistoryHandler) GetTransactions(ctx context.Context, req *v1pb.GetTransactionsRequest) (*v1pb.GetTransactionsResponse, error) {
	txns, err := h.service.GetTransactions(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch transactions: %v", err)
	}
	return &v1pb.GetTransactionsResponse{Transactions: txns}, nil
}
