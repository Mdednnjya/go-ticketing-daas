package repository

import (
	"core-ticketing-engine/internal/entity"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db:db}
}

func (r *TransactionRepository) CreateTransaction(tx *sqlx.Tx, txn entity.Transaction) error {
	query := "INSERT INTO transactions (ticket_id, buyer_email, quantity, total_price) VALUES ($1, $2, $3, $4)"
	_, err := tx.Exec(query, txn.TicketID, txn.BuyerEmail, txn.Quantity, txn.TotalPrice)
	return err

}