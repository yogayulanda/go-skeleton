package model

import "time"

// User adalah model yang digunakan untuk representasi data user
type User struct {
	ID        uint      `gorm:"primaryKey"`      // ID user, primary key
	Name      string    `gorm:"not null"`        // Nama user
	Email     string    `gorm:"not null;unique"` // Email, harus unik
	Password  string    `gorm:"not null"`        // Password yang terenkripsi
	CreatedAt time.Time // Timestamp ketika user dibuat
	UpdatedAt time.Time // Timestamp ketika user diperbarui
}
