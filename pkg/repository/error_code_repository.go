package repository

import (
	"context"
	"errors"
	"time"

	"github.com/yogayulanda/go-skeleton/pkg/models"
	"gorm.io/gorm"
)

type ErrorCodeRepository struct {
	DB *gorm.DB
}

func NewErrorCodeRepository(db *gorm.DB) *ErrorCodeRepository {
	return &ErrorCodeRepository{
		DB: db,
	}
}

// GetAll mengambil semua error codes dari database menggunakan model
func (r *ErrorCodeRepository) GetAll(ctx context.Context) ([]models.ErrorCode, error) {
	var errorCodes []models.ErrorCode

	// Query untuk mengambil semua data error code dari tabel
	if err := r.DB.Find(&errorCodes).Error; err != nil {
		return nil, err
	}

	return errorCodes, nil
}

// GetErrorCode mengambil error code berdasarkan kode dari database
func (r *ErrorCodeRepository) GetErrorMessageByCode(code string) (*models.ErrorCode, error) {
	var errorCode models.ErrorCode
	if err := r.DB.First(&errorCode, "code = ?", code).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Kode error tidak ditemukan, kita bisa mengembalikan error kustom
			return nil, errors.New("error code not found")
		}
		return nil, err
	}
	return &errorCode, nil
}

// GetUpdatedErrorCodes mengambil error codes yang diubah setelah timestamp tertentu
func (r *ErrorCodeRepository) GetUpdatedErrorCodes(ctx context.Context, since time.Time) ([]models.ErrorCode, error) {
	var errorCodes []models.ErrorCode
	err := r.DB.WithContext(ctx).Where("updated_at > ?", since).Find(&errorCodes).Error
	if err != nil {
		return nil, err
	}
	return errorCodes, nil
}
