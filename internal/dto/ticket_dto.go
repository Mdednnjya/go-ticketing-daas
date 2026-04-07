package dto

type TicketResponse struct {
	ID 			int		`json:"id"`
	EventName	string	`json:"event_name"`
	Price		int		`json:"price"`
}