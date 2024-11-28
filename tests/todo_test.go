package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"todolist/config"
	"todolist/models"
	"todolist/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SetupApp() *fiber.App {
	app := fiber.New()

	// Inisialisasi database
	config.ConnectDB()

	// Migrasi database
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})

	// Inisialisasi rute
	routes.AuthRoutes(app)
	routes.TodoRoutes(app)
	return app
}

func getAuthToken(app *fiber.App) (string, error) {
	// Data login
	loginData := map[string]string{
		"username": "testuser1",
		"password": "testpassword",
	}
	jsonLoginData, _ := json.Marshal(loginData)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonLoginData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1) // Nonaktifkan timeout

	if err != nil {
		return "", err
	}

	var response map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	token, exists := response["token"]
	if !exists {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

func TestGetTodos(t *testing.T) {
	app := SetupApp()
	// Mendapatkan token otentikasi
	token, err := getAuthToken(app)
	if err != nil {
		t.Fatalf("Failed to get auth token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateTodo(t *testing.T) {
	app := SetupApp()
	// Mendapatkan token otentikasi
	token, err := getAuthToken(app)
	if err != nil {
		t.Fatalf("Failed to get auth token: %v", err)
	}

	todo := map[string]string{
		"title":       "Test Todo2",
		"description": "Test Description2",
	}
	jsonTodo, _ := json.Marshal(todo)

	req := httptest.NewRequest(http.MethodPost, "/api/todos", bytes.NewBuffer(jsonTodo))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
