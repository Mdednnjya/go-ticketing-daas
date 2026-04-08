package service

import (
	"core-ticketing-engine/internal/dto"
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

func (s *TicketService) GetAllTickets() ([]dto.TicketResponse,error) {
	
	tickets, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get tickets from repository: %w", err)
	}

	var ticketsResponse []dto.TicketResponse

	for _, t := range tickets {
		item := dto.TicketResponse{
			ID: t.ID,
			EventName: t.EventName,
			Price: t.Price,
		}

		ticketsResponse = append(ticketsResponse, item)
	}
	return ticketsResponse, nil
}