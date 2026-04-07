package handler

import (
	"core-ticketing-engine/internal/entity"
	"core-ticketing-engine/internal/service"
	"encoding/json"
	"net/http"
)

type TicketHandler struct {
	service *service.TicketService
}

func NewTicketHandler(service *service.TicketService) *TicketHandler {
	return &TicketHandler{service:service}
}
	

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	// layer 1 Method validation
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error": "Method not Allowed!"}`))
		return
	}

	// layer 2 JSON validation
	var inputTicket entity.Ticket
	err := json.NewDecoder(r.Body).Decode(&inputTicket)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON format!"}`))
		return
	}
	// layer 3 call the service
	err = h.service.CreateTicket(inputTicket)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to process ticket creation"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`"{message": "Ticket successfully created!", "status": "success}"`))
	
}

func (h *TicketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data, err := h.service.GetAllTickets();
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to fetch Data, internal server error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to Encode data"}`))
		return
	}


}