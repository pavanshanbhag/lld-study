package snakeandladdergame

import "sync"

type GameManager struct {
	Games []*SnakeAndLadderGame
}

func NewGameManager() *GameManager {
	return &GameManager{
		Games: []*SnakeAndLadderGame{},
	}
}

// GetGameManager kept for upstream demo compatibility; prefer NewGameManager.

func (gm *GameManager) StartNewGame(wg *sync.WaitGroup, playerNames []string) {
	game := NewSnakeAndLadderGame(playerNames)
	gm.Games = append(gm.Games, game)
	wg.Add(1)
	go game.Play(wg)
}
