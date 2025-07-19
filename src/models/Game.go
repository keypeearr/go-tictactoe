package models

import (
	"errors"
	"log"
	"strings"
	"time"
)

const EMPTY_BOARD = "---------"

var GAME_STATUS = map[string]string{
	"ON-GOING":  "On-Going",
	"COMPLETED": "Completed",
}

var GAME_RESULT = map[string]string{
	"NO RESULT": "",
	"P1 WIN":    "Player 1 Wins",
	"P2 WIN":    "Player 2 Wins",
	"DRAW":      "Draw",
}

/*
Game - Represents a record of the Tic Tac Toe Game
which includes the following: the board where the
players play in, the result of the game, and the
time when the game was played
*/
type Game struct {
	/*
		Board - Stores the data of the tic tac toe in a string format
		Example:
			[O,X,X]
			[0,0,X]
			[X,0,X]
		That matrix becomes "oxxooxxox"
		Where player 1 uses x, while player 2 uses o
	*/
	Board string
	/*
		IsPlayer1Turn - Stores the data of which player will be making a turn
	*/
	IsPlayer1Turn bool
	/*
		Result - Stores the result of the game. There are three possible
		outcomes to this: "Player 1", "Player 2", and "Draw"
	*/
	Result string
	/*
		Status - Stores the status of the game. There are two possible
		outcomes to this: "On-Going", and "Completed"
	*/
	Status string
	/*
		DatePlayed - Stores the time when the game was initiated
	*/
	DatePlayed time.Time
}

func NewGame(isPlayer1First *bool) *Game {
	var cond bool
	if isPlayer1First == nil {
		cond = true
	} else {
		cond = *isPlayer1First
	}

	return &Game{
		Board:         EMPTY_BOARD,
		IsPlayer1Turn: cond,
		Result:        GAME_RESULT["NO RESULT"],
		Status:        GAME_STATUS["ON-GOING"],
		DatePlayed:    time.Now(),
	}
}

func (g *Game) IsBoardFull() bool {
	return !strings.Contains(g.Board, "-")
}

func (g *Game) SwitchTurns() {
	g.IsPlayer1Turn = !g.IsPlayer1Turn
}

func (g *Game) MarkBoard(index int) error {
	if g.IsBoardFull() {
		return errors.New("board is already filled")
	}

	if string(g.Board[index]) != "-" {
		return errors.New("index is already marked")
	}

	mark := "x"
	if g.IsPlayer1Turn == false {
		mark = "o"
	}

	currentBoard := g.Board
	newBoard := currentBoard[:index] + mark + currentBoard[index+1:]
	g.Board = newBoard
	return nil
}

func (g *Game) CheckForWinner() bool {
	comp := "xxx"
	if !g.IsPlayer1Turn {
		comp = "ooo"
	}

	horizontals := []struct {
		Start int
		End   int
	}{
		{0, 2},
		{3, 5},
		{6, 8},
	}
	for _, index := range horizontals {
		str := g.Board[index.Start : index.End+1]
		if str != comp {
			continue
		}
		if comp == "xxx" {
			g.Result = GAME_RESULT["P1 WIN"]
		} else {
			g.Result = GAME_RESULT["P2 WIN"]
		}
		g.Status = GAME_STATUS["COMPLETED"]
		return true
	}

	indexes := []struct {
		Start  int
		Middle int
		End    int
	}{
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}
	for _, index := range indexes {
		str := []byte{
			g.Board[index.Start],
			g.Board[index.Middle],
			g.Board[index.End],
		}
		if string(str) != comp {
			continue
		}
		if comp == "xxx" {
			g.Result = GAME_RESULT["P1 WIN"]
		} else {
			g.Result = GAME_RESULT["P2 WIN"]
		}
		g.Status = GAME_STATUS["COMPLETED"]
		return true
	}

	if g.IsBoardFull() {
		g.Status = GAME_STATUS["COMPLETED"]
		g.Result = GAME_RESULT["DRAW"]
	}
	return false
}

func (g *Game) DisplayBoard() {
	str := ""
	for i, v := range g.Board {
		str += string(v)
		if i == 2 || i == 5 || i == 8 {
			str += "\n"
		}
	}
	log.Println(str)
}

func (g *Game) ResetBoard() {
	g.IsPlayer1Turn = true
	g.Board = EMPTY_BOARD
	g.Status = GAME_STATUS["ON-GOING"]
	g.Result = "NO RESULT"
}
