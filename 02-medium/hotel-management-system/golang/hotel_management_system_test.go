package hotelmanagement

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewHotelManagementSystem()
}
