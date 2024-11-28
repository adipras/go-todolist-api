package controllers

import (
	"strconv"
	"todolist/config"
	"todolist/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetTodos(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["iss"].(string))

	var todos []models.Todo
	config.DB.Where("user_id = ?", userID).Find(&todos)
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID, _ := strconv.Atoi(claims["iss"].(string))

	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	todo.UserID = uint(userID)
	config.DB.Create(&todo)
	return c.JSON(todo)
}

func GetTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
	}
	return c.JSON(todo)
}

func UpdateTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	config.DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodoById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	if err := config.DB.Delete(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Task not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
