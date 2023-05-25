package main

// Importing packages
import (
	"golang-task/initializers"
	"golang-task/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	// LoadEnvVariables()
	initializers.LoadEnvVariables()
	// ConnectToDB()
	initializers.ConnectToDB()
	// SyncDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)

	r.Run() 
}