package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/yogayulanda/go-skeleton/pkg/common"
	"github.com/yogayulanda/go-skeleton/pkg/domain/user"

	"github.com/yogayulanda/go-skeleton/pkg/dto"
	"go.uber.org/zap"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService struct {
	repo user.UserRepository
	log  *zap.Logger
}

func NewUserService(repo user.UserRepository, log *zap.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

// GetUser mengambil data user berdasarkan ID
func (s *UserService) GetUser(ctx context.Context, id string) (*dto.UserDTO, error) {
	// Mengambil user_id dan role dari context menggunakan custom key dari common
	userIDFromCtx, _ := ctx.Value(common.CtxUserID).(string)
	roleFromCtx, _ := ctx.Value(common.CtxRole).(string)

	// Log atau gunakan user_id dan role
	s.log.Info("getting user profile", zap.String("user_id", userIDFromCtx), zap.String("role", roleFromCtx))
	// Konversi string ID ke uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		s.log.Error("Invalid ID format", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	// Mendapatkan user dari repository menggunakan ctx
	user, err := s.repo.GetByID(ctx, uint(userID))
	if err != nil {
		return nil, err
	}

	// Mengonversi Model ke DTO untuk response
	userDTO := &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return userDTO, nil
}

// CreateUser membuat user baru dan mengonversinya ke DTO
func (s *UserService) CreateUser(ctx context.Context, userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	s.log.Info("Creating new user", zap.String("user", userDTO.Name))

	// Mengonversi DTO ke Model
	user := &user.User{
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}

	// Simpan user ke repository
	if err := s.repo.Create(ctx, user); err != nil {
		s.log.Error("Error creating user", zap.String("user", user.Name), zap.Error(err))
		return nil, err
	}

	// Mengonversi Model ke DTO untuk response
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// UpdateUser memperbarui data user berdasarkan ID
func (s *UserService) UpdateUser(ctx context.Context, id string, userDTO *dto.UserDTO) (*dto.UserDTO, error) {
	s.log.Info("Updating user", zap.String("id", id))

	// Konversi string ID ke uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		s.log.Error("Invalid ID format", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	// Mendapatkan user dari repository
	user, err := s.repo.GetByID(ctx, uint(userID))
	if err != nil {
		s.log.Error("Error fetching user", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	// Memperbarui data user
	user.Name = userDTO.Name
	user.Email = userDTO.Email

	// Simpan perubahan ke repository
	if err := s.repo.Update(ctx, user); err != nil {
		s.log.Error("Error updating user", zap.String("id", id), zap.Error(err))
		return nil, err
	}

	// Mengonversi Model ke DTO untuk response
	return &dto.UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// DeleteUser menghapus data user berdasarkan ID
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	s.log.Info("Deleting user", zap.String("id", id))

	// Konversi string ID ke uint
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		s.log.Error("Invalid ID format", zap.String("id", id), zap.Error(err))
		return err
	}

	// Menghapus user dari repository
	if err := s.repo.Delete(ctx, uint(userID)); err != nil {
		s.log.Error("Error deleting user", zap.String("id", id), zap.Error(err))
		return err
	}

	return nil
}
