package main

import (
	"dakbazar/database"
	"dakbazar/internal/models"
	"dakbazar/routes"
	"encoding/json"
	"io/ioutil"
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
		fileData, err := ioutil.ReadFile("data/admins.json")
		if err != nil {
			log.Println("Failed to read file:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Cannot read file"})
		}

		// Step 2: Parse JSON into Go struct
		var admins []models.Admin
		if err := json.Unmarshal(fileData, &admins); err != nil {
			log.Println("Failed to unmarshal JSON:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		// Step 3: Insert into database
		for i := range admins {
			// You can hash password here if needed
			err := database.DB.Create(&admins[i]).Error
			if err != nil {
				log.Println("Failed to insert admin:", err)
				return c.Status(500).JSON(fiber.Map{"error": "Failed to insert admin"})
			}
		}

		return c.JSON(fiber.Map{
			"message": "Data imported successfully!",
		})
	})

	routes.ApiRoutes(app)

	app.Listen(":3000")
}
