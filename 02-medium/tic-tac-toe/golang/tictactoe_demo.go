package tictactoe

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	player1 := NewPlayer("Player 1", 'X')
	player2 := NewPlayer("Player 2", 'O')

	game := NewGame(player1, player2)
	logger.Info("starting game", "p1", player1.Name, "p2", player2.Name)
	game.Play(logger)
}
