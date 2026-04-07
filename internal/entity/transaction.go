package entity

import "time"

type Transaction struct {
	ID         int			`db:"id"`
	TicketID   int			`db:"ticket_id"`
	BuyerEmail string		`db:"buyer_email"`
	Quantity   int			`db:"quantity"`
	TotalPrice int			`db:"total_price"`
	CreateAt   time.Time	`db:"created_at"`
}