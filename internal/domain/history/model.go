package history

type TransactionModel struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
