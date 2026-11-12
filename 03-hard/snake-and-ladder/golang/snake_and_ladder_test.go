package snakeandladdergame

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewSnakeAndLadderGame([]string{"Alice", "Bob"})
	_ = NewGameManager()
}
