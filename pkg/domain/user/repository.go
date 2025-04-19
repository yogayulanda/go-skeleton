package user

import "context"

// UserRepository adalah interface yang mendefinisikan operasi yang dapat dilakukan pada data user
type UserRepository interface {
	GetByID(ctx context.Context, id uint) (*User, error) // Mengambil user berdasarkan ID
	Create(ctx context.Context, user *User) error        // Membuat user baru
	Update(ctx context.Context, user *User) error        // Memperbarui user
	Delete(ctx context.Context, id uint) error           // Menghapus user berdasarkan ID
}
