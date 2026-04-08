package repository

import (
	"core-ticketing-engine/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db:db}
}

func (r *TransactionRepository) CreateTransaction(txn entity.Transaction) error {
	query := "INSERT INTO transactions (ticket_id, buyer_email, quantity, total_price) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(query, txn.TicketID, txn.BuyerEmail, txn.Quantity, txn.TotalPrice)
	if err != nil {
		return fmt.Errorf("Failed during inserting data to db: %w", err)
	}

	return nil

}