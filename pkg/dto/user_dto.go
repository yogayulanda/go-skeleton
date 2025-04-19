package dto

// UserDTO adalah representasi data user yang dikirim/diterima antar layer
type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
