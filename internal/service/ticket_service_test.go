package service

import (
	"errors"
	"testing"

	"core-ticketing-engine/internal/entity" 
)


type mockTicketRepository struct {
	mockFindById func(id int) (entity.Ticket, error)
}

func (m *mockTicketRepository) FindById(id int) (entity.Ticket, error) {
	return m.mockFindById(id)
}

func (m *mockTicketRepository) CreateTicket(t entity.Ticket) error {
	return nil
}

func (m *mockTicketRepository) FindAll() ([]entity.Ticket, error) {
	return nil, nil
}

// scenario
func TestGetTicketByID_Success(t *testing.T) {

	fakeRepo := &mockTicketRepository{
		mockFindById: func(id int) (entity.Ticket, error) {
			return entity.Ticket{
				ID:        1,
				EventName: "Slipknot Jakarta",
				Price:     2500000,
				Stock:     100,
			}, nil
		},
	}

	svc := NewTicketService(fakeRepo, nil)

	result, err := svc.GetTicketByID(1)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if result.EventName != "Slipknot Jakarta" {
		t.Errorf("Expected EventName 'Slipknot Jakarta', got: %v", result.EventName)
	}

	if result.Price != 2500000 {
		t.Errorf("Expected Price 2500000, got: %v", result.Price)
	}
}

// run
func TestGetTicketByID_NotFound(t *testing.T) {
	fakeRepo := &mockTicketRepository{
		mockFindById: func(id int) (entity.Ticket, error) {
			return entity.Ticket{}, errors.New("sql: no rows in result set")
		},
	}

	svc := NewTicketService(fakeRepo, nil)

	result, err := svc.GetTicketByID(99)

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	if result.ID != 0 { 
		t.Errorf("Expected empty ID (0), got: %v", result.ID)
	}
}