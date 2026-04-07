package main

import (
	"core-ticketing-engine/internal/config"
	"core-ticketing-engine/internal/handler"
	"core-ticketing-engine/internal/repository"
	"core-ticketing-engine/internal/service"
	"log"
	"net/http"
)

func main() {
	// init
	mux := http.NewServeMux()
	port := ":8081"
	db := config.ConnectDB()

	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)
	

	// endpoints
	mux.HandleFunc("/health", handler.HealthCheckHandler)
	mux.HandleFunc("/api/tickets", ticketHandler.CreateTicket)

	log.Printf("Server is Running on Port %s \n", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("Error occur %s", err)
	}
}