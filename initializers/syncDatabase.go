package initializers

import (
	"github.com/jwt-project/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
