package routes

import (
	"blockChain/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Block Chain Server is Up and Running")
	})
	app.Post("/addBlock", h.AddBlock)
	app.Post("/findBlock", h.FindBlock)
}
