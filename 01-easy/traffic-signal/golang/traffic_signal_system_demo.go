package trafficsignalsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	trafficController := NewTrafficController()

	road1 := NewRoad("R1", "Main Street")
	road2 := NewRoad("R2", "Broadway")
	road3 := NewRoad("R3", "Park Avenue")
	road4 := NewRoad("R4", "Elm Street")

	road1.SetTrafficLight(NewTrafficLight("TL1", 100, 50, 100))
	road2.SetTrafficLight(NewTrafficLight("TL2", 100, 50, 100))
	road3.SetTrafficLight(NewTrafficLight("TL3", 100, 50, 100))
	road4.SetTrafficLight(NewTrafficLight("TL4", 100, 50, 100))

	trafficController.AddRoad(road1)
	trafficController.AddRoad(road2)
	trafficController.AddRoad(road3)
	trafficController.AddRoad(road4)
	logger.Info("roads registered", "count", trafficController.RoadCount())

	stop := make(chan struct{})
	go trafficController.StartTrafficControl(stop)

	time.Sleep(250 * time.Millisecond)
	if err := trafficController.HandleEmergency("R2"); err != nil {
		logger.Error("emergency failed", "err", err)
	} else {
		logger.Info("emergency green", "road", "R2", "signal", road2.TrafficLight.Signal())
	}
	close(stop)
}
