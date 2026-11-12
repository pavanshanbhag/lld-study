package tictactoe

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Game struct {
	Player1       *Player
	Player2       *Player
	Board         *Board
	CurrentPlayer *Player
}

func NewGame(player1, player2 *Player) *Game {
	return &Game{
		Player1:       player1,
		Player2:       player2,
		Board:         NewBoard(),
		CurrentPlayer: player1,
	}
}

func (g *Game) Play(logger *slog.Logger) {
	if logger == nil {
		logger = slog.Default()
	}
	g.Board.PrintBoard()

	for !g.Board.IsFull() && !g.Board.HasWinner() {
		logger.Info("turn", "player", g.CurrentPlayer.Name)
		row := g.getValidInput("Enter row (0-2): ")
		col := g.getValidInput("Enter column (0-2): ")

		err := g.Board.MakeMove(row, col, g.CurrentPlayer.Symbol)
		if err != nil {
			logger.Error("invalid move", "err", err)
			continue
		}

		g.Board.PrintBoard()
		g.switchPlayer()
	}

	if g.Board.HasWinner() {
		g.switchPlayer()
		logger.Info("winner", "player", g.CurrentPlayer.Name)
	} else {
		logger.Info("draw")
	}
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == g.Player1 {
		g.CurrentPlayer = g.Player2
	} else {
		g.CurrentPlayer = g.Player1
	}
}

func (g *Game) getValidInput(prompt string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input >= 0 && input <= 2 {
			return input
		}
		fmt.Println("Invalid input! Please enter a number between 0 and 2.")
	}
}
