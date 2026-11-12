package airlinemanagementsystem

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewAirlineManagementSystem()
}
