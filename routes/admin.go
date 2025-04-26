package routes

import "github.com/gofiber/fiber/v2"

func AdminRoutes(router fiber.Router) {
	admin := router.Group("/admin")

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from admin route!",
		})
	})
}
