package movieticketbookingsystem

import "sync"

type Booking struct {
	ID         string
	User       *User
	Show       *Show
	Seats      []*Seat
	TotalPrice float64
	status     BookingStatus
	mu         sync.RWMutex
}

func NewBooking(id string, user *User, show *Show, seats []*Seat, totalPrice float64, status BookingStatus) *Booking {
	return &Booking{
		ID:         id,
		User:       user,
		Show:       show,
		Seats:      seats,
		TotalPrice: totalPrice,
		status:     status,
	}
}

func (b *Booking) Status() BookingStatus {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.status
}

func (b *Booking) SetStatus(status BookingStatus) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.status = status
}

