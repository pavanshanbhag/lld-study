package trafficsignalsystem

import "testing"

func TestNewTrafficController(t *testing.T) {
	t.Parallel()
	tc := NewTrafficController()
	road := NewRoad("R1", "Main Street")
	road.SetTrafficLight(NewTrafficLight("L1", 1000, 500, 2000))
	tc.AddRoad(road)
	if len(tc.roads) != 1 {
		t.Fatal("expected one road")
	}
}
