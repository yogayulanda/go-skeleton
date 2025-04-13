package history

import (
	"context"
	"time"

	v1pb "gitlab.twprisma.com/fin/lmd/services/if-trx-history/api/proto/gen/v1"
)

// TrxHistoryService defines the domain interface
type TrxHistoryService interface {
	GetTransactions(ctx context.Context, userID string) ([]*v1pb.Transaction, error)
	StoreTransaction(ctx context.Context, tx *v1pb.Transaction) error
}

type trxHistoryService struct {
	// bisa tambahkan repo kalau pakai database
	// repo Repository
}

// NewTrxHistoryService returns a new service instance
func NewTrxHistoryService() TrxHistoryService {
	return &trxHistoryService{}
}

// GetTransactions fetches transactions (dummy logic for now)
func (s *trxHistoryService) GetTransactions(ctx context.Context, userID string) ([]*v1pb.Transaction, error) {
	txns := []*v1pb.Transaction{
		{
			Id:          "1",
			UserId:      userID,
			Amount:      15000,
			Description: "Topup",
			Timestamp:   time.Now().Format(time.RFC3339),
		},
	}
	return txns, nil
}

// StoreTransaction saves a transaction (dummy logic)
func (s *trxHistoryService) StoreTransaction(ctx context.Context, tx *v1pb.Transaction) error {
	// simpan ke database kalau ada
	return nil
}
