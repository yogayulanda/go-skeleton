package history

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type TransactionRepository struct {
// 	DB *sql.DB
// }

// func NewTransactionRepository(dbURL string) *TransactionRepository {
// 	db, err := sql.Open("mysql", dbURL)
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}
// 	return &TransactionRepository{DB: db}
// }

// func (r *TransactionRepository) FetchTransactionHistory(userID string) ([]repository.Transaction, error) {
// 	var transactions []repository.Transaction
// 	rows, err := r.DB.Query("SELECT * FROM transactions WHERE user_id = ?", userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var t repository.Transaction
// 		if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.CreatedAt); err != nil {
// 			return nil, err
// 		}
// 		transactions = append(transactions, t)
// 	}
// 	return transactions, nil
// }
