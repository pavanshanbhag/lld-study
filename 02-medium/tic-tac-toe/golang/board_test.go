package tictactoe

import "testing"

func TestBoardMakeMoveAndWin(t *testing.T) {
	t.Parallel()

	board := NewBoard()
	playerX := 'X'
	playerO := 'O'

	moves := [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}}
	symbols := []rune{playerX, playerO, playerX, playerO, playerX}

	for i, move := range moves {
		if err := board.MakeMove(move[0], move[1], symbols[i]); err != nil {
			t.Fatalf("MakeMove(%d, %d): %v", move[0], move[1], err)
		}
	}

	if !board.HasWinner() {
		t.Fatal("expected X to win on top row")
	}
	if board.IsFull() {
		t.Fatal("board should not be full yet")
	}
}

func TestBoardInvalidMove(t *testing.T) {
	t.Parallel()

	board := NewBoard()
	if err := board.MakeMove(0, 0, 'X'); err != nil {
		t.Fatalf("first move: %v", err)
	}
	if err := board.MakeMove(0, 0, 'O'); err == nil {
		t.Fatal("expected error for occupied cell")
	}
}

func TestBoardDraw(t *testing.T) {
	t.Parallel()

	board := NewBoard()
	sequence := []struct {
		row, col int
		symbol   rune
	}{
		{0, 0, 'X'}, {0, 1, 'O'}, {0, 2, 'X'},
		{1, 1, 'O'}, {1, 0, 'X'}, {1, 2, 'O'},
		{2, 1, 'X'}, {2, 0, 'O'}, {2, 2, 'X'},
	}

	for _, move := range sequence {
		if err := board.MakeMove(move.row, move.col, move.symbol); err != nil {
			t.Fatalf("MakeMove: %v", err)
		}
	}

	if board.HasWinner() {
		t.Fatal("expected draw, got winner")
	}
	if !board.IsFull() {
		t.Fatal("expected full board")
	}
}

func TestNewGame(t *testing.T) {
	t.Parallel()

	p1 := NewPlayer("Alice", 'X')
	p2 := NewPlayer("Bob", 'O')
	game := NewGame(p1, p2)

	if game.CurrentPlayer != p1 {
		t.Fatal("player 1 should start")
	}
	if game.Board == nil {
		t.Fatal("expected board")
	}
}
