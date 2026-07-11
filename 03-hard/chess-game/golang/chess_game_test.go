package chessgame

import "testing"

func TestChessGameConstructor(t *testing.T) {
	t.Parallel()

	game := NewChessGame()
	if game.board == nil || len(game.players) != 2 {
		t.Fatal("expected board and two players to be initialized")
	}
}

func TestChessGameValidPawnMove(t *testing.T) {
	t.Parallel()

	game := NewChessGame()
	player := game.players[0]
	piece := game.board.GetPiece(1, 0)
	move := NewMove(piece, 2, 0)

	if err := player.MakeMove(game.board, move); err != nil {
		t.Fatalf("MakeMove: %v", err)
	}
	if game.board.GetPiece(2, 0) == nil {
		t.Fatal("expected pawn at destination")
	}
}
