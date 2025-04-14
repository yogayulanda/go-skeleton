package user

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

func (r *SQLRepository) Create(ctx context.Context, user *UserModel) (*UserModel, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SQLRepository) FindByEmail(ctx context.Context, email string) (*UserModel, error) {
	var user UserModel
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SQLRepository) FindByID(ctx context.Context, id string) (*UserModel, error) {
	var user UserModel
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *SQLRepository) Update(ctx context.Context, user *UserModel) (*UserModel, error) {
	if err := r.db.WithContext(ctx).
		Model(&UserModel{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *SQLRepository) Delete(ctx context.Context, id string) (bool, error) {
	result := r.db.WithContext(ctx).Delete(&UserModel{}, "id = ?", id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
