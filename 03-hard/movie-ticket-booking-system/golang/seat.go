package movieticketbookingsystem

import "sync"

type Seat struct {
	ID     string
	Row    int
	Column int
	Type   SeatType
	price  float64
	status SeatStatus
	mu     sync.RWMutex
}

func NewSeat(id string, row, column int, seatType SeatType, price float64, status SeatStatus) *Seat {
	return &Seat{
		ID:     id,
		Row:    row,
		Column: column,
		Type:   seatType,
		price:  price,
		status: status,
	}
}

func (s *Seat) Status() SeatStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.status
}

func (s *Seat) SetStatus(status SeatStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.status = status
}

func (s *Seat) Price() float64 {
	return s.price
}

