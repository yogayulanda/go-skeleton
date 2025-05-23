package handler

import (
	"context"
	"errors"
	"fmt"

	v1 "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/common"
	"github.com/yogayulanda/go-skeleton/pkg/dto" // Menggunakan dto.UserDTO
	"github.com/yogayulanda/go-skeleton/pkg/repository"
	"github.com/yogayulanda/go-skeleton/pkg/service"
	"github.com/yogayulanda/go-skeleton/pkg/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserHandler mengimplementasikan gRPC UserServiceServer
type UserHandler struct {
	v1.UnimplementedUserServiceServer                      // Embedkan UnimplementedUserServiceServer untuk kompatibilitas gRPC
	userService                       *service.UserService // Gunakan pointer ke UserService
	errorRepo                         *repository.ErrorCodeRepository
	log                               *zap.Logger
}

func NewUserHandler(userService *service.UserService, errorRepo *repository.ErrorCodeRepository, log *zap.Logger) *UserHandler { // Menerima pointer
	return &UserHandler{
		userService: userService,
		errorRepo:   errorRepo,
		log:         log,
	}
}

// GetUser menangani permintaan untuk mendapatkan informasi user
func (h *UserHandler) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	h.log.Info("Handling GetUser request", zap.String("id", req.GetId()))

	if req.Id == "" {
		// Gunakan error key dari common
		return nil, status.Errorf(codes.InvalidArgument, common.ErrUserIDRequired)
	}
	// Memanggil UserService untuk mendapatkan user berdasarkan ID
	userDTO, err := h.userService.GetUser(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			// Mengembalikan error not found jika user tidak ditemukan
			return nil, status.Errorf(codes.NotFound, common.ErrUserNotFound)
		}

		//  TODO Move to Utils for get DB and Cache, parameter Ctx, Error, and Code
		errorCode, err := h.errorRepo.GetErrorMessageByCode("E1004") // Contoh menggunakan kode error "E001"
		// Ambil pesan error dari DB berdasarkan kode error
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error fetching error message from DB: %v", err)
		}
		// Gunakan SetGrpcError untuk memberikan response error yang kustom
		return nil, utils.SetGrpcError(codes.Internal, errorCode.Code, errorCode.Message)
	}

	// Mengonversi ID yang bertipe uint ke string
	userID := fmt.Sprintf("%d", userDTO.ID)

	// Mengembalikan response yang sesuai dengan struktur GetUserResponse
	return &v1.GetUserResponse{
		User: &v1.User{
			Id:    userID, // Menggunakan userID yang sudah di-convert ke string
			Name:  userDTO.Name,
			Email: userDTO.Email,
		},
	}, nil
}

// CreateUser menangani permintaan untuk membuat user baru
func (h *UserHandler) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	h.log.Info("Handling CreateUser request", zap.String("user", req.GetUser().GetName()))

	// Mengonversi v1.User dari request gRPC menjadi dto.UserDTO
	userDTO := &dto.UserDTO{
		Name:  req.GetUser().GetName(),
		Email: req.GetUser().GetEmail(),
	}

	// Memanggil UserService untuk membuat user baru
	createdUserDTO, err := h.userService.CreateUser(ctx, userDTO)
	if err != nil {
		h.log.Error("Error creating user", zap.String("user", req.GetUser().GetName()), zap.Error(err))
		return nil, err
	}

	// Mengonversi ID yang bertipe uint ke string
	userID := fmt.Sprintf("%d", createdUserDTO.ID)

	// Mengembalikan response yang sesuai
	return &v1.CreateUserResponse{
		User: &v1.User{
			Id:    userID, // Menggunakan userID yang sudah di-convert ke string
			Name:  createdUserDTO.Name,
			Email: createdUserDTO.Email,
		},
	}, nil
}

// UpdateUser menangani permintaan untuk memperbarui informasi user
func (h *UserHandler) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	h.log.Info("Handling UpdateUser request", zap.String("id", req.GetId()))

	// Mengonversi v1.User dari request gRPC menjadi dto.UserDTO
	userDTO := &dto.UserDTO{
		Name:  req.GetUser().GetName(),
		Email: req.GetUser().GetEmail(),
	}

	// Memanggil UserService untuk memperbarui user
	updatedUserDTO, err := h.userService.UpdateUser(ctx, req.GetId(), userDTO)
	if err != nil {
		h.log.Error("Error updating user", zap.String("id", req.GetId()), zap.Error(err))
		return nil, err
	}

	// Mengonversi ID yang bertipe uint ke string
	userID := fmt.Sprintf("%d", updatedUserDTO.ID)

	// Mengembalikan response yang sesuai
	return &v1.UpdateUserResponse{
		User: &v1.User{
			Id:    userID, // Menggunakan userID yang sudah di-convert ke string
			Name:  updatedUserDTO.Name,
			Email: updatedUserDTO.Email,
		},
	}, nil
}

// DeleteUser menangani permintaan untuk menghapus user
func (h *UserHandler) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	h.log.Info("Handling DeleteUser request", zap.String("id", req.GetId()))

	// Memanggil UserService untuk menghapus user berdasarkan ID
	err := h.userService.DeleteUser(ctx, req.GetId())
	if err != nil {
		h.log.Error("Error deleting user", zap.String("id", req.GetId()), zap.Error(err))
		return nil, err
	}

	// Mengembalikan response yang sesuai
	return &v1.DeleteUserResponse{
		Success: true, // Mengindikasikan bahwa penghapusan berhasil
	}, nil
}
