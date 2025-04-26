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

	routes.SeetingRoutes(app)
	routes.ApiRoutes(app)

	app.Listen(":3000")
}
