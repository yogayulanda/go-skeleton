package handler

import (
	"context"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/gen/proto/v1"
	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/history"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TrxHistoryHandler struct {
	v1pb.UnimplementedTrxHistoryServiceServer
	service history.TrxHistoryService
}

func NewTrxHistoryHandler(service history.TrxHistoryService) *TrxHistoryHandler {
	return &TrxHistoryHandler{service: service}
}

func (h *TrxHistoryHandler) GetTransactions(ctx context.Context, req *v1pb.GetTransactionsRequest) (*v1pb.GetTransactionsResponse, error) {
	txns, err := h.service.GetTransactions(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch transactions: %v", err)
	}
	return &v1pb.GetTransactionsResponse{Transactions: txns}, nil
}

func (h *TrxHistoryHandler) StoreTransaction(ctx context.Context, tx *v1pb.Transaction) (*v1pb.StoreTransactionResponse, error) {
	err := h.service.StoreTransaction(ctx, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store transaction: %v", err)
	}
	return &v1pb.StoreTransactionResponse{StatusMessage: "Transaction stored successfully"}, nil
}
