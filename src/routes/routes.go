package routes

import (
	"github.com/gofiber/fiber/v3"

	gamehandlers "github.com/keypeearr/tictactoe/src/handlers/gameHandlers"
)

func Load(app *fiber.App) {
	views := app.Group("")
	views.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Redirect().Status(fiber.StatusPermanentRedirect).To("/tictactoe")
	})
	views.Get("/tictactoe", gamehandlers.DisplayTictactoe)

	api := app.Group("/api/v1/tictactoe")
	api.Post("/start", gamehandlers.HandleStartGame)
	api.Post("/mark/:index", gamehandlers.HandleMarkBox)
	api.Post("/menu", gamehandlers.HandleMenu)
	api.Post("/reset", gamehandlers.HandleResetGame)
	api.Post("/new", gamehandlers.HandleNewGame)
	api.Post("/continue", gamehandlers.HandleContinueGame)
}
