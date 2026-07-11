package ridesharingservice

import "testing"

func TestRideServiceAddPassengerAndDriver(t *testing.T) {
	t.Parallel()

	service := NewRideService()
	passenger := &Passenger{ID: 1, Name: "Alice", Contact: "555-0100"}
	driver := &Driver{
		ID:       1,
		Name:     "Bob",
		Location: &Location{Latitude: 40.7, Longitude: -74.0},
		Status:   Available,
	}

	service.AddPassenger(passenger)
	service.AddDriver(driver)

	if service.passengers[1] == nil {
		t.Fatal("expected passenger to be registered")
	}
	if service.drivers[1] == nil {
		t.Fatal("expected driver to be registered")
	}
}

func TestRideServiceAcceptRide(t *testing.T) {
	t.Parallel()

	service := NewRideService()
	driver := &Driver{ID: 1, Name: "Bob", Status: Available}
	service.AddDriver(driver)

	ride := &Ride{
		ID:          100,
		Passenger:   &Passenger{ID: 1, Name: "Alice"},
		Source:      &Location{Latitude: 40.7, Longitude: -74.0},
		Destination: &Location{Latitude: 40.8, Longitude: -73.9},
		Status:      Requested,
	}

	service.AcceptRide(driver, ride)
	if ride.Status != Accepted {
		t.Fatalf("ride status = %v, want Accepted", ride.Status)
	}
}
