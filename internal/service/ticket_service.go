package service

import (
	"core-ticketing-engine/internal/entity"
	"core-ticketing-engine/internal/repository"
	"errors"
	"fmt"
)

type TicketService struct {
	repo *repository.TicketRepository
}

func NewTicketService(repo *repository.TicketRepository) *TicketService {
	return &TicketService{repo:repo}
}

func (s *TicketService) CreateTicket(t entity.Ticket) error {
	if t.EventName == "" {
		return errors.New("event name cannot empty")
	}

	if t.Price <= 0 {
		return errors.New("price must be greater then zero")
	}

	err := s.repo.CreateTicket(t)
	if err != nil {
		return fmt.Errorf("failed to create ticket in database: %w", err)
	} 

	return nil

}