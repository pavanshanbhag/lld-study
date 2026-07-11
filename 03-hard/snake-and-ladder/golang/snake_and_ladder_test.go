package snakeandladdergame

import "testing"

func TestGameManagerConstructor(t *testing.T) {
	t.Parallel()

	manager := NewGameManager()
	if manager.Games == nil {
		t.Fatal("expected games slice to be initialized")
	}
}

func TestSnakeAndLadderGameConstructor(t *testing.T) {
	t.Parallel()

	game := NewSnakeAndLadderGame([]string{"Alice", "Bob"})
	if len(game.Players) != 2 || game.Board == nil {
		t.Fatal("expected two players and a board")
	}
}

func TestBoardGetNewPosition(t *testing.T) {
	t.Parallel()

	board := NewBoard()
	if pos := board.GetNewPosition(1); pos != 38 {
		t.Fatalf("ladder at 1 = %d, want 38", pos)
	}
	if pos := board.GetNewPosition(16); pos != 6 {
		t.Fatalf("snake at 16 = %d, want 6", pos)
	}
}
