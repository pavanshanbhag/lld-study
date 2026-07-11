package parkinglot

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	lot := NewParkingLot()
	lot.AddLevel(NewLevel(1, 100))
	lot.AddLevel(NewLevel(2, 80))

	car := NewCar("ABC123")
	truck := NewTruck("XYZ789")
	motorcycle := NewMotorcycle("M1234")

	for _, v := range []Vehicle{car, truck, motorcycle} {
		result, err := lot.Park(v)
		if err != nil {
			logger.Error("park failed", "plate", v.LicensePlate(), "err", err)
			continue
		}
		logger.Info("parked", "plate", v.LicensePlate(), "level", result.LevelFloor, "spot", result.Spot.Number())
	}

	logAvailability(logger, lot, "after parking")

	if _, err := lot.Unpark(motorcycle.LicensePlate()); err != nil {
		logger.Error("unpark failed", "plate", motorcycle.LicensePlate(), "err", err)
	}

	logAvailability(logger, lot, "after unpark")
}

func logAvailability(logger *slog.Logger, lot *ParkingLot, label string) {
	logger.Info("availability snapshot", "label", label, "available", lot.AvailableCount())
	for _, s := range lot.Availability() {
		state := "available"
		if s.Occupied {
			state = "occupied"
		}
		logger.Info("spot", "level", s.Level, "number", s.Number, "type", s.Type, "state", state)
	}
}
