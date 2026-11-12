package hotelmanagement

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	hotelSystem := NewHotelManagementSystem()

	guest1 := NewGuest("G001", "John Doe", "john@example.com", "1234567890")
	guest2 := NewGuest("G002", "Jane Smith", "jane@example.com", "9876543210")
	hotelSystem.AddGuest(guest1)
	hotelSystem.AddGuest(guest2)

	room1 := NewRoom("R001", RoomTypeSingle, 100.0)
	room2 := NewRoom("R002", RoomTypeDouble, 200.0)
	hotelSystem.AddRoom(room1)
	hotelSystem.AddRoom(room2)

	checkInDate := time.Now()
	checkOutDate := checkInDate.AddDate(0, 0, 3)

	reservation1, err := hotelSystem.BookRoom(guest1, room1, checkInDate, checkOutDate)
	if err != nil {
		logger.Error("book failed", "err", err)
		return
	}
	logger.Info("reservation created", "id", reservation1.ID)

	if err := hotelSystem.CheckIn(reservation1.ID); err != nil {
		logger.Error("check-in failed", "err", err)
		return
	}
	logger.Info("checked in", "id", reservation1.ID)

	if err := hotelSystem.CheckOut(reservation1.ID, NewCreditCardPayment()); err != nil {
		logger.Error("check-out failed", "err", err)
		return
	}
	logger.Info("checked out", "id", reservation1.ID)

	if err := hotelSystem.CancelReservation(reservation1.ID); err != nil {
		logger.Info("cancel after checkout (expected)", "err", err)
		return
	}
	logger.Info("reservation cancelled", "id", reservation1.ID)
}
