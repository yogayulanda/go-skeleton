package history

import (
	"context"

	v1pb "github.com/yogayulanda/if-trx-history/gen/proto/v1"
)

type TrxHistoryService struct {
	repository Repository
}

// NewTrxHistoryService returns a new service instance
func NewTrxHistoryService(repo Repository) *TrxHistoryService {
	return &TrxHistoryService{
		repository: repo,
	}
}

// GetTransactions fetches transactions (dummy logic for now)
func (s *TrxHistoryService) GetTransactions(ctx context.Context, userID string) ([]*v1pb.Transaction, error) {
	txns := []*v1pb.Transaction{
		{
			Id:          "1",
			UserId:      userID,
			Amount:      15000,
			Description: "Topup",
		},
	}
	return txns, nil
}

// StoreTransaction saves a transaction (dummy logic)
func (s *TrxHistoryService) StoreTransaction(ctx context.Context, tx *v1pb.Transaction) error {
	// simpan ke database kalau ada
	return nil
}
