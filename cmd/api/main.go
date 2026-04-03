package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"


	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("System is up & Ready!")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("System is up & Ready!"))
	})

	fmt.Printf("Server is Running on Port %s", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		fmt.Printf("Error occur %s", err)
	}
}