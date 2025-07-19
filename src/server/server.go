package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"

	gamehandlers "github.com/keypeearr/tictactoe/src/handlers/gameHandlers"
	"github.com/keypeearr/tictactoe/src/models"
	"github.com/keypeearr/tictactoe/src/routes"
)

func Run() error {
	app := fiber.New(fiber.Config{
		AppName: "Tic Tac Toe",
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use("/", static.New("./src/public"))

	gamehandlers.Engine = models.NewGameEngine()
	routes.Load(app)

	port := "1337"
	serverPort := fmt.Sprintf(":%s", port)
	return app.Listen(serverPort)
}
