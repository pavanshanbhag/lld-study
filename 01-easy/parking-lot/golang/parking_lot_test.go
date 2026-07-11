package parkinglot

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestParkingLotParkAndUnpark(t *testing.T) {
	t.Parallel()

	lot := NewParkingLot()
	lot.AddLevel(NewLevel(1, 10))

	car := NewCar("ABC123")
	result, err := lot.Park(car)
	if err != nil {
		t.Fatalf("park car: %v", err)
	}
	if result.Spot == nil {
		t.Fatal("expected spot")
	}
	if result.LevelFloor != 1 {
		t.Fatalf("level = %d, want 1", result.LevelFloor)
	}

	if lot.AvailableCount() == 10 {
		t.Fatal("expected at least one occupied spot")
	}

	if _, err := lot.Unpark("ABC123"); err != nil {
		t.Fatalf("unpark car: %v", err)
	}

	if _, err := lot.Unpark("ABC123"); !errors.Is(err, ErrVehicleNotFound) {
		t.Fatalf("second unpark err = %v, want ErrVehicleNotFound", err)
	}
}

func TestParkingLotFull(t *testing.T) {
	t.Parallel()

	lot := NewParkingLot()
	lot.AddLevel(NewLevel(1, 2)) // 1 motorcycle spot + 1 truck spot
	if _, err := lot.Park(NewMotorcycle("M1")); err != nil {
		t.Fatalf("park first motorcycle: %v", err)
	}

	if _, err := lot.Park(NewMotorcycle("M2")); !errors.Is(err, ErrLotFull) {
		t.Fatalf("park second motorcycle err = %v, want ErrLotFull", err)
	}
}

func TestParkingLotConcurrentPark(t *testing.T) {
	t.Parallel()

	lot := NewParkingLot()
	lot.AddLevel(NewLevel(1, 50))

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			_, _ = lot.Park(NewCar(fmt.Sprintf("CAR-%d", n)))
		}(i)
	}
	wg.Wait()

	if lot.AvailableCount() >= 50 {
		t.Fatal("expected some spots occupied")
	}
}
