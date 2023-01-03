package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", ListFacts)

	app.Post("/fact", CreateFact)
}
