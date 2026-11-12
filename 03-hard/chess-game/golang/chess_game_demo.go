package chessgame

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	logger.Info("starting chess game demo")
	game := NewChessGame()
	game.Start()
}
