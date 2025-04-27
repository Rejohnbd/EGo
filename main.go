package main

import (
	"dakbazar/database"
	"dakbazar/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	database.ConnectMysqlDB()

	app := fiber.New(fiber.Config{
		ReadBufferSize: 4096 * 10,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SeetingRoutes(app)
	routes.ApiRoutes(app)

	appURL := os.Getenv("APP_URL")
	appPort := os.Getenv("APP_PORT")

	if appURL == "" {
		appURL = "0.0.0.0"
	}
	if appPort == "" {
		appPort = "3000"
	}

	listenAddress := appURL + ":" + appPort

	log.Println("Server running at http://" + listenAddress)

	app.Listen(listenAddress)
}
