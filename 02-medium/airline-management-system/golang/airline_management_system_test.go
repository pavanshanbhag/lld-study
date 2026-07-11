package airlinemanagementsystem

import (
	"testing"
	"time"
)

func TestAirlineManagementSystemBookAndCancel(t *testing.T) {
	t.Parallel()

	system := NewAirlineManagementSystem()
	passenger := NewPassenger("P1", "Alice", "alice@example.com", "1234567890")
	departure := time.Now().Add(24 * time.Hour)
	flight := NewFlight("F001", "New York", "London", departure, departure.Add(2*time.Hour))
	system.AddFlight(flight)

	seat := NewSeat("12A", SeatTypeEconomy)
	booking := system.BookFlight(flight, passenger, seat, 250.0)
	if booking == nil || booking.BookingNumber == "" {
		t.Fatal("expected non-empty booking")
	}

	system.CancelBooking(booking.BookingNumber)
	if booking.Status != BookingStatusCancelled {
		t.Fatalf("booking status = %v, want cancelled", booking.Status)
	}
}
