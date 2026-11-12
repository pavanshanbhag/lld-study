package chessgame

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewChessGame()
}
