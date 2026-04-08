package repository

import (
	"core-ticketing-engine/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TicketRepository struct {
	db *sqlx.DB
}

func NewTicketRepository(db *sqlx.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) CreateTicket(ticket entity.Ticket) error {
	query := "INSERT INTO tickets (event_name, price) VALUES ($1, $2)"
	
	_, err := r.db.Exec(query, ticket.EventName, ticket.Price)

	if err != nil {
		return fmt.Errorf("Error inserting data to db: %w", err)
	}

	return nil
}

func (r *TicketRepository) FindAll() ([]entity.Ticket, error) {
	query := "SELECT * FROM tickets"
	var tickets []entity.Ticket

	err := r.db.Select(&tickets, query)

	if err != nil {
		return nil, fmt.Errorf("Failed to Fetch ticket from db: %w", err)
	}

	return tickets, nil
}

func (r *TicketRepository) FindByID(id int) (entity.Ticket, error) {
	query := "SELECT * FROM tickets WHERE ID = $1"
	var ticket entity.Ticket
	
	err := r.db.Get(&ticket, query, id)
	if err != nil {
		return entity.Ticket{}, fmt.Errorf("Failed to fetch ticket from db: %w", err)
	}

	return ticket, nil
}