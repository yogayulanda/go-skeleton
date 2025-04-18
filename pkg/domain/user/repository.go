package user

import (
	"context"
)

// Repository adalah interface yang mendefinisikan operasi database untuk entitas User
type Repository interface {
	// Create menyimpan pengguna baru
	Create(ctx context.Context, user *UserModel) (*UserModel, error)

	// FindByEmail mencari pengguna berdasarkan email
	FindByEmail(ctx context.Context, email string) (*UserModel, error)

	// FindByID mencari pengguna berdasarkan ID
	FindByID(ctx context.Context, id string) (*UserModel, error)

	// Update memperbarui informasi pengguna
	Update(ctx context.Context, user *UserModel) (*UserModel, error)

	// Delete menghapus pengguna berdasarkan ID
	Delete(ctx context.Context, id string) (bool, error)
}
