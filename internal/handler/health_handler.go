package handler

import (
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("System is up and Running!\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System is running and health!"))
}