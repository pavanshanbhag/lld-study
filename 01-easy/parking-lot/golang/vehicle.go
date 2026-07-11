package parkinglot

import "fmt"

type VehicleType int

const (
	VehicleMotorcycle VehicleType = iota
	VehicleCar
	VehicleTruck
)

func (t VehicleType) String() string {
	switch t {
	case VehicleMotorcycle:
		return "motorcycle"
	case VehicleCar:
		return "car"
	case VehicleTruck:
		return "truck"
	default:
		return fmt.Sprintf("VehicleType(%d)", t)
	}
}

type Vehicle interface {
	LicensePlate() string
	Type() VehicleType
}

type vehicle struct {
	licensePlate string
	vehicleType  VehicleType
}

func (v vehicle) LicensePlate() string { return v.licensePlate }
func (v vehicle) Type() VehicleType    { return v.vehicleType }

func NewCar(licensePlate string) Vehicle {
	return vehicle{licensePlate: licensePlate, vehicleType: VehicleCar}
}

func NewMotorcycle(licensePlate string) Vehicle {
	return vehicle{licensePlate: licensePlate, vehicleType: VehicleMotorcycle}
}

func NewTruck(licensePlate string) Vehicle {
	return vehicle{licensePlate: licensePlate, vehicleType: VehicleTruck}
}
