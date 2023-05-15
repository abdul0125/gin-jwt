package initializers

import "gin-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
