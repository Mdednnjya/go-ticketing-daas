package service

import (
	"core-ticketing-engine/internal/dto"
	"core-ticketing-engine/internal/entity"
	"core-ticketing-engine/internal/repository"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	ticketRepo *repository.TicketRepository
	db *sqlx.DB
}

func NewTransactionService(TxnRepo *repository.TransactionRepository, TicRepo *repository.TicketRepository, db *sqlx.DB) *TransactionService {
	return &TransactionService{transactionRepo: TxnRepo, ticketRepo: TicRepo, db: db}
}


func (s *TransactionService) CreateTransaction(req dto.TransactionRequest) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed tp init tx object: %w", err)
		
	}

	// unit of work 
	ticket, err := s.ticketRepo.FindByIDTx(tx, req.TicketID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find ticket: %w", err)
	}

	if req.Quantity > ticket.Stock {
		tx.Rollback()
		return errors.New("empty stock")
	}

	err = s.ticketRepo.DecreaseStock(tx, req.TicketID, req.Quantity)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to decrease stock: %w", err)
	}

	data := entity.Transaction{
		TicketID: req.TicketID,
		BuyerEmail: req.BuyerEmail,
		Quantity: req.Quantity,
		TotalPrice: req.Quantity * ticket.Price,
	}

	err = s.transactionRepo.CreateTransaction(tx, data)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert transaction into db: %w", err)
	}
	
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}