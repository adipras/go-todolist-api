package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/config"
	"todolist/models"
	"todolist/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SetupAuthApp() *fiber.App {
	app := fiber.New()

	// Inisialisasi database
	config.ConnectDB()

	// Migrasi database
	config.DB.AutoMigrate(&models.User{})

	// Inisialisasi rute
	routes.AuthRoutes(app)
	return app
}

func TestRegister(t *testing.T) {
	app := SetupAuthApp()

	user := map[string]string{
		"username": "testuser1",
		"password": "testpassword",
	}
	jsonUser, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseUser models.User
	json.NewDecoder(resp.Body).Decode(&responseUser)
	assert.Equal(t, "testuser1", responseUser.Username)
}

func TestLogin(t *testing.T) {
	app := SetupAuthApp()

	// Buat pengguna terlebih dahulu
	user := models.User{
		Username: "testuser2",
		Password: "$2a$14$UAZM7ZyQPz7ib7cCsrTwVerjjeT6ZBt.RV/7ZM0fujju4ul1k2deW", // "testpassword" hashed
	}
	config.DB.Create(&user)

	loginData := map[string]string{
		"username": "testuser2",
		"password": "testpassword",
	}
	jsonLoginData, _ := json.Marshal(loginData)

	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(jsonLoginData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	_, exists := response["token"]
	assert.True(t, exists)
}
