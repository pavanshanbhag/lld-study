package elevatorsystem

import "testing"

func TestElevatorControllerRequest(t *testing.T) {
	t.Parallel()

	controller := NewElevatorController(2, 5)
	defer controller.Stop()

	if err := controller.RequestElevator(1, 5); err != nil {
		t.Fatalf("RequestElevator(1,5): %v", err)
	}
	if err := controller.RequestElevator(3, 7); err != nil {
		t.Fatalf("RequestElevator(3,7): %v", err)
	}
}

func TestElevatorAddRequestAtCapacity(t *testing.T) {
	t.Parallel()

	elevator := NewElevator(1, 1)
	elevator.Run()
	defer elevator.Stop()

	if err := elevator.AddRequest(NewRequest(1, 2)); err != nil {
		t.Fatalf("first request: %v", err)
	}
	if err := elevator.AddRequest(NewRequest(2, 3)); err == nil {
		t.Fatal("expected capacity error")
	}
}
