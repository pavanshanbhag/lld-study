package airlinemanagementsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	system := NewAirlineManagementSystem()

	passenger1 := NewPassenger("U001", "John Doe", "john@example.com", "1234567890")

	departureTime1 := time.Now().AddDate(0, 0, 1)
	arrivalTime1 := departureTime1.Add(2 * time.Hour)
	flight1 := NewFlight("F001", "New York", "London", departureTime1, arrivalTime1)

	departureTime2 := time.Now().AddDate(0, 0, 3)
	arrivalTime2 := departureTime2.Add(5 * time.Hour)
	flight2 := NewFlight("F002", "Paris", "Tokyo", departureTime2, arrivalTime2)

	system.AddFlight(flight1)
	system.AddFlight(flight2)
	system.AddAircraft(NewAircraft("A001", "Boeing 747", 300))
	system.AddAircraft(NewAircraft("A002", "Airbus A380", 500))

	searchDate := time.Now().AddDate(0, 0, 1)
	for _, flight := range system.SearchFlights("New York", "London", searchDate) {
		logger.Info("flight found", "number", flight.FlightNumber, "from", flight.Source, "to", flight.Destination)
	}

	seat := NewSeat("25A", SeatTypeEconomy)
	booking := system.BookFlight(flight1, passenger1, seat, 100.0)
	if booking == nil {
		logger.Error("booking failed")
		return
	}
	logger.Info("booking successful", "id", booking.BookingNumber)

	system.CancelBooking(booking.BookingNumber)
	logger.Info("booking cancelled", "id", booking.BookingNumber)
}
