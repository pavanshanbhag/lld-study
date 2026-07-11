package carrentalsystem

import (
	"testing"
	"time"
)

func TestRentalSystemMakeAndCancelReservation(t *testing.T) {
	t.Parallel()

	rs := NewRentalSystem()
	car := NewCar("Toyota", "Camry", 2022, "ABC123", 50.0)
	rs.AddCar(car)

	customer := NewCustomer("Alice", "alice@example.com", "DL1234")
	start := time.Now()
	end := start.AddDate(0, 0, 2)

	reservation, err := rs.MakeReservation(customer, car, start, end)
	if err != nil {
		t.Fatalf("MakeReservation: %v", err)
	}
	if car.IsAvailable() {
		t.Fatal("car should be unavailable after reservation")
	}

	rs.CancelReservation(reservation.ReservationID)
	if !car.IsAvailable() {
		t.Fatal("car should be available after cancellation")
	}
}
