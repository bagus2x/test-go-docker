package main

import (
	"github.com/bagus2x/database"
	"github.com/bagus2x/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	handlers.SetupRoutes(app)

	app.Listen(":3000")
}
