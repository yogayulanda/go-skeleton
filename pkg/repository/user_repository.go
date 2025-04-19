package repository

import (
	"context"

	"github.com/yogayulanda/go-skeleton/pkg/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db  *gorm.DB
	log *zap.Logger
}

// NewUserRepository adalah konstruktor untuk UserRepository
func NewUserRepository(db *gorm.DB, log *zap.Logger) *UserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

// GetByID mengambil user berdasarkan ID, dengan ctx
func (r *UserRepository) GetByID(ctx context.Context, id uint) (*user.User, error) {
	var u user.User
	// Menggunakan ctx di dalam query
	if err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error; err != nil {
		r.log.Error("Error fetching user by ID", zap.Uint("id", id), zap.Error(err))
		return nil, err
	}
	return &u, nil
}

// Create menyimpan user baru ke dalam database
func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		r.log.Error("Error creating user", zap.String("user", u.Name), zap.Error(err))
		return err
	}
	return nil
}

// Update memperbarui data user
func (r *UserRepository) Update(ctx context.Context, u *user.User) error {
	if err := r.db.WithContext(ctx).Save(u).Error; err != nil {
		r.log.Error("Error updating user", zap.String("user", u.Name), zap.Error(err))
		return err
	}
	return nil
}

// Delete menghapus user berdasarkan ID
func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&user.User{}, "id = ?", id).Error; err != nil {
		r.log.Error("Error deleting user", zap.Uint("id", id), zap.Error(err))
		return err
	}
	return nil
}
