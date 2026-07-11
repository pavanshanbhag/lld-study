package parkinglot

import "errors"

var (
	ErrLotFull          = errors.New("no available spot for vehicle")
	ErrVehicleNotFound  = errors.New("vehicle not parked in lot")
	ErrInvalidVehicle   = errors.New("invalid vehicle")
	ErrSpotNotAvailable = errors.New("spot not available")
)
