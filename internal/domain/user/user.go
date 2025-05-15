package user

// User adalah model yang digunakan untuk representasi data user
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}
