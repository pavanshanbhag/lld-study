package elevatorsystem

import "testing"

func TestElevatorControllerRequest(t *testing.T) {
	t.Parallel()

	controller := NewElevatorController(2, 5)
	defer controller.Stop()

	controller.RequestElevator(1, 5)
	controller.RequestElevator(3, 7)
}
