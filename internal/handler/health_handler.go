package handler

import (
	"fmt"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("System is up and Running!\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("System is running and health!"))
}