package handler

import (
	"core-ticketing-engine/internal/entity"
	"net/http"	
	"encoding/json"
)
	

func TicketsHandler(w http.ResponseWriter, r *http.Request) {
	ticket01 := []entity.Ticket{
		{
			ID: 1, 
			EventName: "Coldplay",
			Price: 150,
		},
		{
			ID: 2,
			EventName: "Coldplay JKT",
			Price: 175,
		},
	}

	data, err := json.Marshal(ticket01)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{error: Internal Server Error!}"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	
}