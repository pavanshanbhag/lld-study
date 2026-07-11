package parkinglot

type Level struct {
	floor int
	spots []*ParkingSpot
}

func NewLevel(floor, numSpots int) *Level {
	level := &Level{floor: floor}

	bikeSpots := int(float64(numSpots) * 0.50)
	carSpots := int(float64(numSpots) * 0.40)

	for i := 1; i <= bikeSpots; i++ {
		level.spots = append(level.spots, NewParkingSpot(i, VehicleMotorcycle))
	}
	for i := bikeSpots + 1; i <= bikeSpots+carSpots; i++ {
		level.spots = append(level.spots, NewParkingSpot(i, VehicleCar))
	}
	for i := bikeSpots + carSpots + 1; i <= numSpots; i++ {
		level.spots = append(level.spots, NewParkingSpot(i, VehicleTruck))
	}

	return level
}

func (l *Level) Floor() int { return l.floor }

func (l *Level) Park(v Vehicle) (*ParkingSpot, error) {
	for _, spot := range l.spots {
		if spot.Available() && spot.VehicleType() == v.Type() {
			if err := spot.Park(v); err != nil {
				continue
			}
			return spot, nil
		}
	}
	return nil, ErrLotFull
}

func (l *Level) Unpark(licensePlate string) (*ParkingSpot, error) {
	for _, spot := range l.spots {
		if spot.Available() {
			continue
		}
		if spot.ParkedVehicle().LicensePlate() == licensePlate {
			if _, err := spot.Unpark(); err != nil {
				return nil, err
			}
			return spot, nil
		}
	}
	return nil, ErrVehicleNotFound
}

func (l *Level) Availability() []SpotStatus {
	status := make([]SpotStatus, 0, len(l.spots))
	for _, spot := range l.spots {
		status = append(status, spot.Status(l.floor))
	}
	return status
}
