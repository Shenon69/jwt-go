package initializers

import "github.com/Shenon69/jwt-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}