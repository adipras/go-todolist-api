package routes

import (
	"todolist/controllers"
	"todolist/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app *fiber.App) {
	api := app.Group("/api/todos", middlewares.Auth())

	api.Get("/", controllers.GetTodos)
	api.Get("/:id", controllers.GetTodoById)
	api.Post("/", controllers.CreateTodo)
	api.Put("/:id", controllers.UpdateTodoById)
	api.Delete("/:id", controllers.DeleteTodoById)
}
