package main

import (
	"log"
	"todolist/config"
	"todolist/migrations"
	"todolist/routes"
	"todolist/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	config.ConnectDB()
	migrations.Migrate()

	utils.SetupLogger()

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Listen(":3000")
}

//
