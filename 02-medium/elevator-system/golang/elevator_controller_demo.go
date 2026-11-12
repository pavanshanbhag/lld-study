package elevatorsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	controller := NewElevatorController(3, 5)
	defer controller.Stop()

	for _, req := range [][2]int{{5, 10}, {3, 7}, {8, 2}, {1, 9}} {
		if err := controller.RequestElevator(req[0], req[1]); err != nil {
			logger.Error("request failed", "from", req[0], "to", req[1], "err", err)
			continue
		}
		logger.Info("elevator requested", "from", req[0], "to", req[1])
	}

	time.Sleep(500 * time.Millisecond)
	logger.Info("demo complete")
}
