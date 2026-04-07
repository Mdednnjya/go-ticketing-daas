package dto

type TransactionRequest struct {
	TicketID	int		`json:"ticket_id"`
	BuyerEmail	string	`json:"buyer_email"`
	Quantity	int		`json:"quantity"`

}