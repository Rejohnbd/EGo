package routes

import (
	"dakbazar/database"

	"github.com/gofiber/fiber/v2"
)

func SeetingRoutes(router fiber.Router) {
	router.Get("/import-data", func(c *fiber.Ctx) error {

		if err := database.SeedStatus(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedStatus failed!"})
		}

		if err := database.SeedAdmins(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedAdmins failed!"})
		}

		if err := database.SeedZones(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedZones failed!"})
		}

		if err := database.SeedMediaUploads(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedMediaUploads failed!"})
		}

		if err := database.SeedAdminShippingMethods(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedAdminShippingMethods failed!"})
		}

		return c.JSON(fiber.Map{"message": "Seeding completed!"})
	})
}
