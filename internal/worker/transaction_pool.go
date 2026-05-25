package worker

import (
	"core-ticketing-engine/internal/dto"
	"core-ticketing-engine/internal/service"
)

const (
	WorkerCount = 5   // fixed nums
	QueueSize   = 100 // buffered chan capacity
)

type TransactionJob struct {
	Request dto.TransactionRequest
	Result  chan error
}

type TransactionPool struct {
	service *service.TransactionService
	jobs    chan TransactionJob
}

func NewTransactionPool(svc *service.TransactionService) *TransactionPool {
	return &TransactionPool{
		service: svc,
		jobs:    make(chan TransactionJob, QueueSize),
	}
}

// start worker
func (p *TransactionPool) Start() {
	for i := 0; i < WorkerCount; i++ {
		go p.worker(i)
	}
}

// worker loop
func (p *TransactionPool) worker(id int) {
	for job := range p.jobs {
		err := p.service.CreateTransaction(job.Request)
		job.Result <- err
	}
}

// submit job, wait for result
func (p *TransactionPool) Submit(req dto.TransactionRequest) error {
	job := TransactionJob{
		Request: req,
		Result:  make(chan error, 1),
	}

	// block if queue full
	p.jobs <- job 

	return <-job.Result
}
