package controllers

import (
	"net/http"

	"golang-task/initializers"
	"golang-task/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email/password from request body
	var body struct {
		Email string 
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fields are empty",
		})
		return	
	}


	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	// create user
	user := models.User{ Email: body.Email, Password: string(hash) }
	result := initializers.DB.Create(&user)
	
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error creating user",
		})
		return
	}


	// respond
	c.JSON(http.StatusOK, gin.H{})
}