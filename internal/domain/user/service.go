package user

import (
	"context"
	"errors"
)

// Service adalah struct yang menangani logika bisnis pengguna
type Service struct {
	repo Repository // repository untuk user
}

// NewService untuk membuat instance baru dari Service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateUser membuat pengguna baru
func (s *Service) CreateUser(ctx context.Context, user *UserModel) (*UserModel, error) {
	// Validasi input jika diperlukan
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	// Cek apakah email sudah digunakan
	existingUser, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already used")
	}

	// Simpan pengguna baru
	return s.repo.Create(ctx, user)
}

// GetUser mengembalikan pengguna berdasarkan ID
func (s *Service) GetUser(ctx context.Context, id string) (*UserModel, error) {
	if id == "" {
		return nil, errors.New("id required")
	}
	return s.repo.FindByID(ctx, id)
}

// UpdateUser memperbarui pengguna
func (s *Service) UpdateUser(ctx context.Context, user *UserModel) (*UserModel, error) {
	if user.ID == "" {
		return nil, errors.New("id required")
	}
	// Update pengguna
	return s.repo.Update(ctx, user)
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *Service) DeleteUser(ctx context.Context, id string) (bool, error) {
	if id == "" {
		return false, errors.New("id required")
	}
	// Hapus pengguna
	return s.repo.Delete(ctx, id)
}
