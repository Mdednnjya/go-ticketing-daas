package handler

import (
	"core-ticketing-engine/internal/dto"
	"core-ticketing-engine/internal/worker"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	pool *worker.TransactionPool
}

func NewTransactionHandlder(pool *worker.TransactionPool) *TransactionHandler {
	return &TransactionHandler{pool: pool}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}"error": "invalid JSON format"!}`))
		return
	}

	err = h.pool.Submit(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to proccess transactions"}`))
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status": "success", "message": "transaction succesfully created!"}`))
}