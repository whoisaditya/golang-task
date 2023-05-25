package main

// Importing packages
import (
	"golang-task/controllers"
	"golang-task/initializers"
	"golang-task/middleware"

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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Golang Task",
		})
	})
	
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/addDetails", middleware.RequireAuth, controllers.AddDetails)
	r.Run() 
}