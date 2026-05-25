package main

import (
	"core-ticketing-engine/internal/config"
	"core-ticketing-engine/internal/handler"
	"core-ticketing-engine/internal/repository"
	"core-ticketing-engine/internal/service"
	"core-ticketing-engine/internal/worker"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	// init
	mux  := http.NewServeMux()
	port := ":8081"
	db   := config.ConnectDB()
	rdb  := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// ticket domain
	ticketRepo := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepo, rdb)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// transaction domain
	txnRepo := repository.NewTransactionRepository(db)
	txnService := service.NewTransactionService(txnRepo, ticketRepo, db)
	txnPool := worker.NewTransactionPool(txnService)
	txnPool.Start()
	txnHandler := handler.NewTransactionHandler(txnPool)


	// endpoints
	mux.HandleFunc("/health", handler.HealthCheckHandler)
	mux.HandleFunc("POST /api/tickets", ticketHandler.CreateTicket)
	mux.HandleFunc("GET /api/tickets", ticketHandler.GetTickets)
	mux.HandleFunc("GET /api/ticket/{id}", ticketHandler.GetTicketById)
	mux.HandleFunc("POST /api/transactions", txnHandler.CreateTransaction)

	log.Printf("Server is Running on Port %s \n", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("Error occur %s", err)
	}
}