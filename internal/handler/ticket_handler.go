package handler

import (
	"core-ticketing-engine/internal/entity"
	"net/http"	
	"encoding/json"
)
	

func TicketsHandler(w http.ResponseWriter, r *http.Request) {
	ticket01 := entity.Ticket{
		ID: 01, 
		EventName: "Coldplay",
		Price: 150,
	}

	data, err := json.Marshal(ticket01)

	if err != nil {
		// edge case logic
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	
}