package service

import (
	"context"
	"core-ticketing-engine/internal/dto"
	"core-ticketing-engine/internal/entity"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)


type TicketRepository interface {
	CreateTicket(entity.Ticket) error
	FindAll() ([]entity.Ticket, error)
	FindById(id int) (entity.Ticket, error)
}

type TicketService struct {
	repo TicketRepository
	rdb *redis.Client
}

func NewTicketService(repo TicketRepository, rdb *redis.Client) *TicketService {
	return &TicketService{repo:repo, rdb: rdb}
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
			Stock: t.Stock,
		}

		ticketsResponse = append(ticketsResponse, item)
	}
	return ticketsResponse, nil
}

func (s *TicketService) GetTicketByID(id int) (dto.TicketResponse, error) {
	ctx := context.Background()
	redisKey := fmt.Sprintf("ticket: %d", id)

	// cache hit
	if s.rdb != nil {
		cacheData, err := s.rdb.Get(ctx, redisKey).Result()
		if err == nil {
			log.Println("Cache hit: Retriving ticket data from redis")

			var ticketResponse dto.TicketResponse

			json.Unmarshal([]byte(cacheData), &ticketResponse)

			return ticketResponse, nil
		}
		
	}

	// cache miss
	ticket, err := s.repo.FindById(id)
	log.Println("Cache miss: Retriving ticket data from Psql")
	if err != nil {
		return dto.TicketResponse{}, err
	}

	ticketResponse := dto.TicketResponse {
		ID: ticket.ID,
		EventName: ticket.EventName,
		Price: ticket.Price,
		Stock: ticket.Stock,
	} 

	// sync
	if s.rdb != nil {
		ticketJSON, _ := json.Marshal(ticketResponse)

		// TTL 5 Minute
		s.rdb.Set(ctx, redisKey, ticketJSON, 5*time.Minute)
	}


	return ticketResponse, nil
}