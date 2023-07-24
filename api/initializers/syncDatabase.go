package initializers

import "github.com/whoisaditya/golang-task/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}