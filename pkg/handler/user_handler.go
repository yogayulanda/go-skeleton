package handler

import (
	"context"

	userProto "github.com/yogayulanda/go-skeleton/gen/proto/v1"
	"github.com/yogayulanda/go-skeleton/pkg/domain/user"
	"github.com/yogayulanda/go-skeleton/pkg/utils"
)

type UserHandler struct {
	userProto.UnimplementedUserServiceServer // embed gRPC compatibility
	userService                              *user.Service
}

func NewUserHandler(userService *user.Service) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser untuk mengambil pengguna berdasarkan ID
func (h *UserHandler) GetUser(ctx context.Context, req *userProto.GetUserRequest) (*userProto.GetUserResponse, error) {
	// Memanggil service untuk mengambil user berdasarkan ID
	user, err := h.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	// Mengonversi model user ke proto User
	return &userProto.GetUserResponse{
		User: &userProto.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// CreateUser untuk membuat pengguna baru
func (h *UserHandler) CreateUser(ctx context.Context, req *userProto.CreateUserRequest) (*userProto.CreateUserResponse, error) {
	// Mengambil user dari request dan memetakan ke model User
	user := &user.UserModel{
		ID:    req.GetUser().GetId(),
		Name:  req.GetUser().GetName(),
		Email: req.GetUser().GetEmail(),
	}

	// Panggil service untuk membuat user
	createdUser, err := h.userService.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Mengubah model user ke format proto
	responseProto := &userProto.CreateUserResponse{}
	if err := utils.ModelToProto(createdUser, responseProto); err != nil {
		return nil, err
	}

	// Mengembalikan response setelah sukses membuat user
	return responseProto, nil
}

// UpdateUser untuk memperbarui informasi pengguna
func (h *UserHandler) UpdateUser(ctx context.Context, req *userProto.UpdateUserRequest) (*userProto.UpdateUserResponse, error) {
	// Konversi proto ke model (struktur pkg)
	user := &user.UserModel{}
	if err := utils.ProtoToModel(req.GetUser(), user); err != nil {
		return nil, err
	}

	// Panggil service untuk memperbarui user
	updatedUser, err := h.userService.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Konversi model kembali ke proto untuk respons
	responseProto := &userProto.UpdateUserResponse{}
	if err := utils.ModelToProto(updatedUser, responseProto); err != nil {
		return nil, err
	}

	return responseProto, nil
}

// DeleteUser untuk menghapus pengguna berdasarkan ID
func (h *UserHandler) DeleteUser(ctx context.Context, req *userProto.DeleteUserRequest) (*userProto.DeleteUserResponse, error) {
	// Panggil service untuk menghapus user
	success, err := h.userService.DeleteUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	// Mengembalikan response setelah menghapus user
	return &userProto.DeleteUserResponse{
		Success: success,
	}, nil
}
