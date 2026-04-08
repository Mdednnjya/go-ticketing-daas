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

	// ticket domain
	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// transaction domain
	txnRepo := repository.NewTransactionRepository(db)
	txnService := service.NewTransactionService(txnRepo, ticketRepo)
	txnHandler := handler.NewTransactionHandlder(txnService)

	

	// endpoints
	mux.HandleFunc("/health", handler.HealthCheckHandler)
	mux.HandleFunc("POST /api/tickets", ticketHandler.CreateTicket)
	mux.HandleFunc("GET /api/tickets", ticketHandler.GetTickets)
	mux.HandleFunc("POST /api/transactions", txnHandler.CreateTransaction)

	log.Printf("Server is Running on Port %s \n", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("Error occur %s", err)
	}
}