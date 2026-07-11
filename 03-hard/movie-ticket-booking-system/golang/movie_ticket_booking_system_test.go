package movieticketbookingsystem

import (
	"testing"
	"time"
)

func TestMovieTicketBookingSystemBookTickets(t *testing.T) {
	t.Parallel()

	system := NewMovieTicketBookingSystem()
	movie := NewMovie("M1", "Inception", "Sci-fi thriller", 148)
	theater := NewTheater("T1", "Cineplex", "Downtown")
	seats := CreateSeats(2, 2)
	show := NewShow("S1", movie, theater, time.Now(), time.Now().Add(2*time.Hour), seats)
	user := NewUser("U1", "Alice", "alice@example.com")

	system.AddMovie(movie)
	system.AddTheater(theater)
	system.AddShow(show)

	selectedSeats := []*Seat{seats["1-1"]}
	booking, err := system.BookTickets(user, show, selectedSeats)
	if err != nil {
		t.Fatalf("BookTickets: %v", err)
	}
	if booking == nil || booking.TotalPrice != 150 {
		t.Fatalf("booking total = %v, want 150", booking.TotalPrice)
	}
}
