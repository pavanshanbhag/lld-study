package parkinglot

import (
	"sync"
)

type ParkingLot struct {
	mu     sync.RWMutex
	levels []*Level
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{levels: make([]*Level, 0)}
}

func (p *ParkingLot) AddLevel(level *Level) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.levels = append(p.levels, level)
}

func (p *ParkingLot) Park(v Vehicle) (*ParkingSpot, error) {
	if v == nil || v.LicensePlate() == "" {
		return nil, ErrInvalidVehicle
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, level := range p.levels {
		spot, err := level.Park(v)
		if err == nil {
			return spot, nil
		}
		if err != ErrLotFull {
			return nil, err
		}
	}
	return nil, ErrLotFull
}

func (p *ParkingLot) Unpark(licensePlate string) (*ParkingSpot, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, level := range p.levels {
		spot, err := level.Unpark(licensePlate)
		if err == nil {
			return spot, nil
		}
		if err != ErrVehicleNotFound {
			return nil, err
		}
	}
	return nil, ErrVehicleNotFound
}

func (p *ParkingLot) Availability() []SpotStatus {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var status []SpotStatus
	for _, level := range p.levels {
		status = append(status, level.Availability()...)
	}
	return status
}

func (p *ParkingLot) AvailableCount() int {
	count := 0
	for _, s := range p.Availability() {
		if !s.Occupied {
			count++
		}
	}
	return count
}
