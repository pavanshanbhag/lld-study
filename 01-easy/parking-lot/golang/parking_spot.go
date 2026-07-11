package parkinglot

type ParkingSpot struct {
	number      int
	vehicleType VehicleType
	vehicle     Vehicle
}

func NewParkingSpot(number int, vehicleType VehicleType) *ParkingSpot {
	return &ParkingSpot{number: number, vehicleType: vehicleType}
}

func (s *ParkingSpot) Number() int { return s.number }

func (s *ParkingSpot) VehicleType() VehicleType { return s.vehicleType }

func (s *ParkingSpot) Available() bool { return s.vehicle == nil }

func (s *ParkingSpot) Park(v Vehicle) error {
	if !s.Available() {
		return ErrSpotNotAvailable
	}
	if v.Type() != s.vehicleType {
		return ErrSpotNotAvailable
	}
	s.vehicle = v
	return nil
}

func (s *ParkingSpot) Unpark() (Vehicle, error) {
	if s.Available() {
		return nil, ErrVehicleNotFound
	}
	v := s.vehicle
	s.vehicle = nil
	return v, nil
}

func (s *ParkingSpot) ParkedVehicle() Vehicle { return s.vehicle }

type SpotStatus struct {
	Level    int
	Number   int
	Type     VehicleType
	Occupied bool
}

func (s *ParkingSpot) Status(level int) SpotStatus {
	return SpotStatus{
		Level:    level,
		Number:   s.number,
		Type:     s.vehicleType,
		Occupied: !s.Available(),
	}
}
