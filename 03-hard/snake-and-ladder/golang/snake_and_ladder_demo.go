package snakeandladdergame

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	gameManager := NewGameManager()
	wg := new(sync.WaitGroup)
	// Start game 1
	players1 := []string{"Player 1", "Player 2", "Player 3"}
	gameManager.StartNewGame(wg, players1)

	// Start game 2
	players2 := []string{"Player 4", "Player 5"}
	gameManager.StartNewGame(wg, players2)

	logger.Info("demo", "msg", fmt.Sprint("Games started. Check game output above."))
	//Wait for all games to finish
	wg.Wait()
}
