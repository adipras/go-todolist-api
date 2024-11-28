package migrations

import (
	"todolist/config"
	"todolist/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Todo{})
}
