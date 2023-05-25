package initializers

import "golang-task/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}