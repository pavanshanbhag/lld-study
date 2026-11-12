package carrentalsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	rentalSystem := NewRentalSystem()

	rentalSystem.AddCar(NewCar("Toyota", "Camry", 2022, "ABC123", 50.0))
	rentalSystem.AddCar(NewCar("Honda", "Civic", 2021, "XYZ789", 45.0))
	rentalSystem.AddCar(NewCar("Ford", "Mustang", 2023, "DEF456", 80.0))

	customer := NewCustomer("John Doe", "john@example.com", "DL1234")
	startDate := time.Now()
	endDate := startDate.AddDate(0, 0, 3)

	availableCars := rentalSystem.SearchCars("Toyota", "Camry", startDate, endDate)
	if len(availableCars) == 0 {
		logger.Info("no available cars")
		return
	}

	reservation, err := rentalSystem.MakeReservation(customer, availableCars[0], startDate, endDate)
	if err != nil {
		logger.Error("reservation failed", "err", err)
		return
	}

	if err := rentalSystem.ProcessPayment(reservation); err != nil {
		logger.Error("payment failed", "err", err)
		rentalSystem.CancelReservation(reservation.ReservationID)
		return
	}
	logger.Info("reservation successful", "id", reservation.ReservationID)
}
