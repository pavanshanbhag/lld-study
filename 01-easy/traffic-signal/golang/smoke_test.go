package trafficsignalsystem

import "testing"

func TestNewTrafficController(t *testing.T) {
	t.Parallel()
	tc := NewTrafficController()
	road := NewRoad("R1", "Main Street")
	road.SetTrafficLight(NewTrafficLight("L1", 1000, 500, 2000))
	tc.AddRoad(road)
	if tc.RoadCount() != 1 {
		t.Fatal("expected one road")
	}
}

func TestHandleEmergency(t *testing.T) {
	t.Parallel()
	tc := NewTrafficController()
	road := NewRoad("R1", "Main")
	light := NewTrafficLight("L1", 10, 10, 10)
	road.SetTrafficLight(light)
	tc.AddRoad(road)

	if err := tc.HandleEmergency("R1"); err != nil {
		t.Fatalf("emergency: %v", err)
	}
	if light.Signal() != Green {
		t.Fatalf("signal = %v, want Green", light.Signal())
	}
}

func TestHandleEmergencyMissingRoad(t *testing.T) {
	t.Parallel()
	tc := NewTrafficController()
	if err := tc.HandleEmergency("missing"); err == nil {
		t.Fatal("expected error for missing road")
	}
}
