package initializers

import "github.com/whoisaditya/golang-task/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}