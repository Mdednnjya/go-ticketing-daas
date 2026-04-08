package service

import (
	"core-ticketing-engine/internal/dto"
	"core-ticketing-engine/internal/entity"
	"core-ticketing-engine/internal/repository"
	"fmt"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	ticketRepo *repository.TicketRepository
}

func NewTransactionService(TxnRepo *repository.TransactionRepository, TicRepo *repository.TicketRepository) *TransactionService {
	return &TransactionService{transactionRepo: TxnRepo, ticketRepo: TicRepo}
}


func (s *TransactionService) CreateTransaction(req dto.TransactionRequest) error {
	ticket, err := s.ticketRepo.FindByID(req.TicketID)
	if err != nil {
		return fmt.Errorf("failed to fint ticket: %w", err)
	}

	data := entity.Transaction{
		TicketID: req.TicketID,
		BuyerEmail: req.BuyerEmail,
		Quantity: req.Quantity,
		TotalPrice: req.Quantity * ticket.Price,
	}

	err = s.transactionRepo.CreateTransaction(data)
	if err != nil {
		return fmt.Errorf("failed to insert transaction into db: %w", err)
	}

	return nil
}