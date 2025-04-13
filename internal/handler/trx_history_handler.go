package handler

import (
	"context"
	"time"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
)

// TrxHistoryHandler mengimplementasikan v1pb.TrxHistoryServiceServer
type TrxHistoryHandler struct {
	v1pb.UnimplementedTrxHistoryServiceServer
}

func NewTrxHistoryHandler() *TrxHistoryHandler {
	return &TrxHistoryHandler{}
}

func (h *TrxHistoryHandler) GetTransactions(ctx context.Context, req *v1pb.GetTransactionsRequest) (*v1pb.GetTransactionsResponse, error) {
	// Simulasi ambil data transaksi dari service/domain
	txns := []*v1pb.Transaction{
		{
			Id:          "1",
			UserId:      req.UserId,
			Amount:      15000,
			Description: "Topup",
			Timestamp:   time.Now().Format(time.RFC3339),
		},
	}
	return &v1pb.GetTransactionsResponse{Transactions: txns}, nil
}

func (h *TrxHistoryHandler) StoreTransaction(ctx context.Context, tx *v1pb.Transaction) (*v1pb.StoreTransactionResponse, error) {
	// Simulasi simpan transaksi
	return &v1pb.StoreTransactionResponse{StatusMessage: "Transaction stored successfully"}, nil
}
