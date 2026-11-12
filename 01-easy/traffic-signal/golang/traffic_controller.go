package trafficsignalsystem

import (
	"fmt"
	"sync"
	"time"
)

type TrafficController struct {
	roads map[string]*Road
	mu    sync.Mutex
}

func NewTrafficController() *TrafficController {
	return &TrafficController{roads: make(map[string]*Road)}
}

func (tc *TrafficController) AddRoad(road *Road) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.roads[road.ID] = road
}

func (tc *TrafficController) RoadCount() int {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return len(tc.roads)
}

func (tc *TrafficController) StartTrafficControl(stop <-chan struct{}) {
	tc.mu.Lock()
	roads := make([]*Road, 0, len(tc.roads))
	for _, road := range tc.roads {
		roads = append(roads, road)
	}
	tc.mu.Unlock()

	for _, road := range roads {
		trafficLight := road.TrafficLight
		go func(tl *TrafficLight) {
			for {
				select {
				case <-stop:
					return
				case <-time.After(time.Duration(tl.RedDuration) * time.Millisecond):
				}
				tl.ChangeSignal(Green)
				select {
				case <-stop:
					return
				case <-time.After(time.Duration(tl.GreenDuration) * time.Millisecond):
				}
				tl.ChangeSignal(Yellow)
				select {
				case <-stop:
					return
				case <-time.After(time.Duration(tl.YellowDuration) * time.Millisecond):
				}
				tl.ChangeSignal(Red)
			}
		}(trafficLight)
	}
}

func (tc *TrafficController) HandleEmergency(roadID string) error {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	road, exists := tc.roads[roadID]
	if !exists {
		return fmt.Errorf("road not found: %s", roadID)
	}
	road.TrafficLight.ChangeSignal(Green)
	return nil
}
