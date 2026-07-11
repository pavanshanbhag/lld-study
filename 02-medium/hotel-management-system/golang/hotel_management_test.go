package hotelmanagement

import (
	"testing"
	"time"
)

func TestHotelManagementSystemBookRoom(t *testing.T) {
	t.Parallel()

	system := NewHotelManagementSystem()
	guest := NewGuest("G001", "Alice", "alice@example.com", "1234567890")
	room := NewRoom("R001", RoomTypeSingle, 100.0)
	system.AddGuest(guest)
	system.AddRoom(room)

	checkIn := time.Now()
	checkOut := checkIn.AddDate(0, 0, 2)

	reservation, err := system.BookRoom(guest, room, checkIn, checkOut)
	if err != nil {
		t.Fatalf("BookRoom: %v", err)
	}
	if reservation == nil || room.GetStatus() != RoomStatusBooked {
		t.Fatal("expected booked room and reservation")
	}
}
