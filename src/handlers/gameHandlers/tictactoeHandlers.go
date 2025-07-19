package gamehandlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"github.com/keypeearr/tictactoe/src/models"
	"github.com/keypeearr/tictactoe/src/utils"
	"github.com/keypeearr/tictactoe/src/views/pages"
	"github.com/keypeearr/tictactoe/src/views/props"
)

var Engine *models.GameEngine

func DisplayTictactoe(ctx fiber.Ctx) error {
	p := props.TictactoePageProps{
		Game: Engine.CurrentGame,
		MainLayoutProps: props.MainLayoutProps{
			Title: "Tic Tac Toe",
		},
	}

	return utils.Render(ctx, pages.Tictactoe(p))
}

func HandleStartGame(ctx fiber.Ctx) error {
	if err := Engine.StartNewGame(nil); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return utils.Render(ctx, pages.Error(err))
	}

	return utils.Render(ctx, pages.Game(*Engine.CurrentGame))
}

func HandleMarkBox(ctx fiber.Ctx) error {
	i := ctx.Params("index")
	index, err := strconv.Atoi(i)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return utils.Render(ctx, pages.Error(err))
	}

	if err := Engine.MarkBoard(index); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return utils.Render(ctx, pages.Error(err))
	}

	return utils.Render(ctx, pages.Game(*Engine.CurrentGame))
}

func HandleMenu(ctx fiber.Ctx) error {
	return utils.Render(ctx, pages.MainMenu(Engine.CurrentGame))
}

func HandleResetGame(ctx fiber.Ctx) error {
	Engine.ResetCurrentBoard()

	return utils.Render(ctx, pages.Game(*Engine.CurrentGame))
}

func HandleNewGame(ctx fiber.Ctx) error {
	if err := Engine.StartNewGame(nil); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return utils.Render(ctx, pages.Error(err))
	}

	return utils.Render(ctx, pages.Game(*Engine.CurrentGame))
}

func HandleContinueGame(ctx fiber.Ctx) error {
	if hasGame := Engine.CurrentGame != nil; !hasGame {
		ctx.Status(fiber.StatusInternalServerError)
		err := errors.New("there is no existing game")
		return utils.Render(ctx, pages.Error(err))
	}
	return utils.Render(ctx, pages.Game(*Engine.CurrentGame))
}
