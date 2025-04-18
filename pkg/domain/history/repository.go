package history

import (
	"context"
)

// Repository adalah interface yang mendefinisikan operasi database untuk entitas trxHistory
type Repository interface {
	// Create menyimpan pengguna baru
	Create(ctx context.Context, trxHistory *TransactionModel) (*TransactionModel, error)

	// FindByEmail mencari pengguna berdasarkan email
	FindByEmail(ctx context.Context, email string) (*TransactionModel, error)

	// FindByID mencari pengguna berdasarkan ID
	FindByID(ctx context.Context, id string) (*TransactionModel, error)

	// Update memperbarui informasi pengguna
	Update(ctx context.Context, trxHistory *TransactionModel) (*TransactionModel, error)
}
