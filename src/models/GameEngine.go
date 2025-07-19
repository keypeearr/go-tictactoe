package models

import (
	"errors"
)

type GameEngine struct {
	CurrentGame *Game
	History     []Game
	Message     *string
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		CurrentGame: nil,
		History:     []Game{},
	}
}

func (ge *GameEngine) StartNewGame(isPlayer1First *bool) error {
	if ge.CurrentGame != nil {
		ge.History = append(ge.History, *ge.CurrentGame)
	}

	ge.CurrentGame = NewGame(isPlayer1First)
	msg := "game has begun"
	ge.Message = &msg
	return nil
}

func (ge *GameEngine) MarkBoard(index int) error {
	if ge.CurrentGame.Status == GAME_STATUS["COMPLETED"] {
		return errors.New("game is already completed")
	}

	if err := ge.CurrentGame.MarkBoard(index); err != nil {
		return err
	}

	hasWinner := ge.CurrentGame.CheckForWinner()
	ge.CurrentGame.SwitchTurns()
	if hasWinner {
		msg := ge.GenerateWinMessage()
		ge.Message = &msg
		return nil
	}
	return nil
}

func (ge *GameEngine) GenerateWinMessage() string {
	if ge.CurrentGame.Status != GAME_STATUS["COMPLETED"] {
		return ""
	}

	if ge.CurrentGame.Result == GAME_RESULT["DRAW"] {
		return "the game ended with a draw"
	}

	if ge.CurrentGame.Result == GAME_RESULT["P1 WIN"] {
		return "player 1 wins"
	} else {
		return "player 2 wins"
	}
}

func (ge *GameEngine) ResetCurrentBoard() {
	ge.CurrentGame.ResetBoard()
}
