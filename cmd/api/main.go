package main

import (
	"core-ticketing-engine/internal/handler"
	"log"
	"net/http"
)

func main() {
	// init
	mux := http.NewServeMux()
	port := ":8080"

	// endpoints
	mux.HandleFunc("/health", handler.HealthCheckHandler)
	mux.HandleFunc("/api/tickets", handler.TicketsHandler)

	log.Printf("Server is Running on Port %s \n", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalf("Error occur %s", err)
	}
}