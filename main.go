package main

import (
	"log"
	"todolist/config"
	"todolist/migrations"
	"todolist/routes"
	"todolist/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	config.ConnectDB()
	migrations.Migrate()

	utils.SetupLogger()

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Listen(":3000")
}

//
