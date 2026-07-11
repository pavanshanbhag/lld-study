package concertbookingsystem

import (
	"testing"
	"time"
)

func TestBookingSystemBookAndCancel(t *testing.T) {
	t.Parallel()

	system := NewBookingSystem()
	seats := GenerateSeats(10)
	concert := NewConcert("C001", "Artist", "Venue", time.Now().Add(24*time.Hour), seats)
	system.AddConcert(concert)

	user := NewUser("U001", "Alice", "alice@example.com")
	selectedSeats := concert.Seats[:2]

	booking, err := system.BookTickets(user, concert, selectedSeats)
	if err != nil {
		t.Fatalf("BookTickets: %v", err)
	}
	if booking == nil {
		t.Fatal("expected booking")
	}

	for _, seat := range selectedSeats {
		if seat.GetStatus() != StatusBooked {
			t.Fatalf("seat %s should be booked", seat.SeatNumber)
		}
	}

	system.CancelBooking(booking.ID)
	for _, seat := range selectedSeats {
		if seat.GetStatus() != StatusAvailable {
			t.Fatalf("seat %s should be available after cancel", seat.SeatNumber)
		}
	}
}
