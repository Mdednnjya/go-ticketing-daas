package repository

import (
	"core-ticketing-engine/internal/entity"
	"log"

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
		log.Fatalf("Error inserting data to DB! \n")
		return err
	}

	return nil
}