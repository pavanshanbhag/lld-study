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

func (gm *GameManager) StartNewGame(wg *sync.WaitGroup, playerNames []string) {
	game := NewSnakeAndLadderGame(playerNames)
	gm.Games = append(gm.Games, game)
	wg.Add(1)
	go game.Play(wg)
}
