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
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	query := "INSERT INTO transactions (ticket_id, buyer_email, quantity, total_price) VALUES ($1, $2, $3, $4)"

	_, err = tx.Exec(query, txn.TicketID, txn.BuyerEmail, txn.Quantity, txn.TotalPrice)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed during inserting data to db, rolling back: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil

}