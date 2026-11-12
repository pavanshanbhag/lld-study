package concertbookingsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	bookingSystem := NewBookingSystem()

	concert1 := NewConcert("C001", "Artist 1", "Venue 1", time.Now().Add(30*24*time.Hour), GenerateSeats(100))
	concert2 := NewConcert("C002", "Artist 2", "Venue 2", time.Now().Add(60*24*time.Hour), GenerateSeats(50))
	bookingSystem.AddConcert(concert1)
	bookingSystem.AddConcert(concert2)

	user1 := NewUser("U001", "John Doe", "john@example.com")
	user2 := NewUser("U002", "Jane Smith", "jane@example.com")

	for _, concert := range bookingSystem.SearchConcerts("Artist 1", "Venue 1", time.Now().Add(30*24*time.Hour)) {
		logger.Info("concert found", "artist", concert.Artist, "venue", concert.Venue)
	}

	booking1, err := bookingSystem.BookTickets(user1, concert1, concert1.Seats[:3])
	if err != nil {
		logger.Error("booking failed", "err", err)
	} else {
		logger.Info("booking created", "id", booking1.ID)
	}

	booking2, err := bookingSystem.BookTickets(user2, concert2, concert2.Seats[:2])
	if err != nil {
		logger.Error("booking failed", "err", err)
	} else {
		logger.Info("booking successful", "id", booking2.ID)
	}

	if booking1 != nil {
		bookingSystem.CancelBooking(booking1.ID)
		logger.Info("booking cancelled", "id", booking1.ID)
	}

	booking3, err := bookingSystem.BookTickets(user2, concert1, concert1.Seats[3:5])
	if err != nil {
		logger.Error("booking failed", "err", err)
	} else {
		logger.Info("booking successful", "id", booking3.ID)
	}
}
