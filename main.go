package main

import (
	"dakbazar/database"
	"dakbazar/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/import-data", func(c *fiber.Ctx) error {

		if err := database.SeedStatus(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedStatus failed!"})
		}

		if err := database.SeedAdmins(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedAdmins failed!"})
		}

		if err := database.SeedZones(); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "SeedZones failed!"})
		}

		return c.JSON(fiber.Map{"message": "Seeding completed!"})
	})

	routes.ApiRoutes(app)

	app.Listen(":3000")
}
