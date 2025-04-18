package history

import (
	"context"

	"gorm.io/gorm"
)

type SQLRepository struct {
	db *gorm.DB
}

func NewSQLRepository(db *gorm.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

// Pastikan implementasi sesuai interface
var _ Repository = (*SQLRepository)(nil)

func (r *SQLRepository) Create(ctx context.Context, user *TransactionModel) (*TransactionModel, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SQLRepository) FindByEmail(ctx context.Context, email string) (*TransactionModel, error) {
	var user TransactionModel
	return &user, nil
}

func (r *SQLRepository) FindByID(ctx context.Context, id string) (*TransactionModel, error) {
	var user TransactionModel
	return &user, nil
}

func (r *SQLRepository) Update(ctx context.Context, user *TransactionModel) (*TransactionModel, error) {
	return user, nil
}
