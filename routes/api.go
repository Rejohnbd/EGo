package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from API route!",
		})
	})

	AdminRoutes(api)
	VendorRoutes(api)
}
