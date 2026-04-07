package entity

type Ticket struct {
	ID 			int 	`db:"id"`
	EventName 	string	`db:"event_name"`
	Price 		int		`db:"price"`

}