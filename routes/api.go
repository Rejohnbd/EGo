package routes

import (
	"dakbazar/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello from API route!",
		})
	})

	api.Get("/sliders", controllers.Sliders)
	api.Get("/featured-items", controllers.FeaturedItems)

	AdminRoutes(api)
	VendorRoutes(api)
}
