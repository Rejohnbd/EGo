package routes

import "github.com/gofiber/fiber/v2"

func VendorRoutes(router fiber.Router) {
	vendor := router.Group("/vendor")

	vendor.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from vendor route!",
		})
	})
}
