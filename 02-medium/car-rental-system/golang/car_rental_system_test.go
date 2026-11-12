package carrentalsystem

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewRentalSystem()
}
